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

	"github.com/golang/glog"
	rpsl "github.com/manrs-tools/contrib/rpsl-parser"
	rppb "github.com/manrs-tools/contrib/rpsl-parser/proto"
)

var (
	threads = flag.Int("threads", 4, "Max threads to use in parsing db files.")
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
func parseFile(ic <-chan string, rc chan<- *rppb.Record, ec chan<- bool) {
	for fn := range ic {
		var rdr *rpsl.Reader

		reader, err := getReader(fn)
		if err != nil {
			glog.Infof("Failed to read file(%v): %v\n", fn, err)
			continue
		}
		rdr = rpsl.NewReader(reader)

		// Read all leading comments and whitespace.
		err = rdr.ConsumeComment()
		if err != nil {
			glog.Infof("Failed reading file(%v): %v\n", fn, err)
			continue
		}
		err = rdr.ConsumeLeadingWS()
		if err != nil {
			glog.Infof("Failed reading file(%v): %v\n", fn, err)
			continue
		}

		// The file must start with a letter, all IRR records start with a letter character.
		r := rdr.Peek()
		if !rpsl.IsLetter(r) {
			glog.Infof("The first character read(%v) is not a letter, file unparsable.\n", string(r))
			// Add 2 more chars so finding the problem is more possible.
			r, _, _ := rdr.Read()
			glog.Infof("Next char: %v\n", string(r))
			r, _, _ = rdr.Read()
			glog.Infof("Next char: %v\n", string(r))
			continue
		}

		// Parse the file, sending results back up the channel (rc).
		rpsl.Parse(rdr, rc)

		// When done parsing, send a file-done message on the end channel (ec).
		ec <- true
	}
}

// Verify that there files requested exist, open each in a goroutine and feed
// those to the rpsl-parser library, returning each record to a channel for
// disposition in the final data structure to be loaded into a DB.
func main() {
	flag.Var(&rpslFiles, "rpslFiles", "Files to parse, irr/rpsl content, filenames as csv.")
	flag.Parse()

	// Flag checks.
	if len(rpslFiles) == 0 {
		flag.PrintDefaults()
		return
	}

	// Two buffered channels, one for input and for the resulting records.
	ic := make(chan string, len(rpslFiles))
	rc := make(chan *rppb.Record, 100)

	// To signal that all files are done processing,
	// use ec (end channel) to pass state of 'done with file X'
	// to the Record processor.
	ec := make(chan bool, len(rpslFiles))

	// Start the parsing/worker thread
	for i := 0; i < *threads; i++ {
		go parseFile(ic, rc, ec)
	}

	// Push each file into the input (ic) channel.
	for _, fn := range rpslFiles {
		ic <- fn
	}
	close(ic)

	// Track the number of files completed.
	files := 0
	// Read records from the channel,
	// TODO(morrowc): I'm positive this is supposed to be simpler with
	// a sync.WaitGroup. Investigate that later.
Loop:
	for {
		select {
		case r := <-rc:
			fmt.Printf("Record returned: %v\n", r)
			if files == len(rpslFiles) && len(rc) == 0 {
				break Loop
			}
		case <-ec:
			files++
			if files == len(rpslFiles) && len(rc) == 0 {
				break Loop
			}
		default:
		}
	}

	close(rc)
}
