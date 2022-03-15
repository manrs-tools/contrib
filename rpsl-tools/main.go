//
// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// irrParser loads the content of IRR database backups, parses the content
// with the rpsl-parser library and generates loadable data for indexed
// queries of the irr content.
package main

import (
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"unicode"

	"github.com/golang/glog"
	"github.com/manrs-tools/contrib/rpsl-parser"
)

var (
	threads = flag.Int("threads", 4, "Max threads to use in parsing db files.")
	display = flag.Bool("display", false, "Display decoded records on screen")
)

// Define a slice of filenames to get as a single flag on startup.
type fileSlice []string

var rpslFiles fileSlice

// Stirng and Set are required for the variadic flag: rpslFiles.
func (f *fileSlice) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *fileSlice) Set(value string) error {
	for _, fn := range strings.Split(value, ",") {
		*f = append(*f, fn)
	}
	return nil
}

// getReader parses the filename and either returns a gzip reader or regular reader.
func getReader(fn string) (io.Reader, error) {
	fd, err := os.Open(fn)
	if err != nil {
		fmt.Printf("Failed to open file(%v): %v\n", fn, err)
		return nil, err
	}
	if strings.HasSuffix(fn, ".gz") {
		f, err := gzip.NewReader(fd)
		if err != nil {
			fmt.Printf("Failed to read gzip'd file(%v): %v\n", fn, err)
			return nil, err
		}
		return f, nil
	}
	return fd, nil
}

// parseFile reads and parses files as their filenames arrive on
// the input channel. A parse error will abort processing for the
// corresponding file and move to the next one
func parseFile(fn string, sem chan int, records *rpsl.Records, wg *sync.WaitGroup) {
	defer wg.Done()
	sem <- 1

	var rdr *rpsl.Reader

	reader, err := getReader(fn)
	if err != nil {
		glog.Infof("Failed to read file(%v): %v\n", fn, err)
		<-sem
		return
	}
	rdr = rpsl.NewReader(reader)

	// Read all leading comments and whitespace.
	err = rdr.ConsumeComment()
	if err != nil {
		glog.Infof("Failed reading file(%v): %v\n", fn, err)
		<-sem
		return
	}
	err = rdr.ConsumeLeadingWS()
	if err != nil {
		glog.Infof("Failed reading file(%v): %v\n", fn, err)
		<-sem
		return
	}

	// The file must start with a letter, all IRR records start with a letter character.
	r := rdr.Peek()
	if !unicode.IsLetter(r) {
		glog.Infof("The first character read(%v) is not a letter, file unparsable.\n", string(r))
		// Add 2 more chars so finding the problem is more possible.
		r, _, _ := rdr.Read()
		glog.Infof("Next char: %v\n", string(r))
		r, _, _ = rdr.Read()
		glog.Infof("Next char: %v\n", string(r))
		<-sem
		return
	}

	// Parse the file, sending results back up the channel (rc).
	rpsl.Parse(rdr, records)
	<-sem
}

// Verify that the files requested exist, open each in a goroutine and feed
// those to the rpsl-parser library, returning each record to a channel for
// disposition in the final data structure to be loaded into a DB.
func main() {
	//defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	flag.Var(&rpslFiles, "rpslFiles", "Files to parse, irr/rpsl content, filenames as csv.")
	flag.Parse()

	// Flag checks.
	if len(rpslFiles) == 0 {
		flag.PrintDefaults()
		return
	}
	numFiles := len(rpslFiles)
	// If the amount of file is less than threads, only spin up enough threads for those files
	if numFiles < *threads {
		*threads = numFiles
	}

	var records rpsl.Records

	sem := make(chan int, *threads)
	var wg sync.WaitGroup
	wg.Add(numFiles)
	for _, fn := range rpslFiles {
		go parseFile(fn, sem, &records, &wg)
	}
	wg.Wait()

	fmt.Printf("Received a total of %d records\n", len(records.Records))
}
