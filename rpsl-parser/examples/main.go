package main

import (
	"fmt"
	"os"

	rpsl "github.com/manrs-tools/contrib/rpsl-parser"
	rppb "github.com/manrs-tools/contrib/rpsl-parser/proto"
)

const (
	irrFile = "/tmp/radb.db"
)

func main() {
	// Open the irr file, creating an io.Reader for the rpsl.Reader{}.
	fd, err := os.Open(irrFile)
	if err != nil {
		fmt.Printf("Failed to open irrFile(%v): %v\n", irrFile, err)
		return
	}

	// Make a new rpsl.Reader{}.
	rdr := rpsl.NewReader(fd)

	// Read the first character from the file.
	// Use this character to validate that the file MIGHT be correct.
	r, _, err := rdr.Read()
	if err != nil {
		fmt.Printf("failed to readRune: %v\n", err)
		return
	}
	err = rdr.Unread()
	if err != nil {
		fmt.Printf("failed to unRead a rune(%v): %v\n", r, err)
		return
	}

	// The file must start with a letter, all IRR records start with a letter character.
	if !rpsl.IsLetter(r) {
		fmt.Printf("the initial character read(%v) is not a letter, file unparsable.", r)
		return
	}

	// Start to Parse() the reader contents, report results up the channel (rc).
	rc := make(chan *rppb.Record)
	go rpsl.Parse(rdr, rc)

	for rec := range rc {
		fmt.Printf("Record type(%v):\n", rec.Type)
		// Ideally this is where actual processing of the Record{} based upon
		// type, content, etc happens. For this example code the record is
		// simply output to stdout.
		for k, v := range rec.Fields {
			fmt.Printf("Key(%v)\t-> Val(%v)\n", k, v)
		}
		fmt.Println()
	}
}
