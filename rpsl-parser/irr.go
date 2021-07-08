//
// Copyright 2018 Google LLC
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

// Package irr implements a parser for IRR/RPSL data.
package rpsl

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"

	rppb "github.com/manrs-tools/contrib/rpsl-parser/proto"

	glog "github.com/golang/glog"
)

// Reader is a struct to manage access to the irr database file.
type Reader struct {
	reader *bufio.Reader
}

// NewReader instantiates a new reader object.
func NewReader(fd io.Reader) *Reader {
	return &Reader{reader: bufio.NewReader(fd)}
}

// Read exports the ReadRune() method for the bufio.Reader.
func (r *Reader) Read() (rune, int, error) {
	return r.reader.ReadRune()
}

// Unread exports the UnreadRune() method for the bufio.Reader.
func (r *Reader) Unread() error {
	return r.reader.UnreadRune()
}

// findKey scans forward until a colon is discovered, returning the keyword found.
func (r *Reader) findKey() (rppb.Type, string) {
	var buf bytes.Buffer
	for {
		// ReadRule until a colon is found, return the buf and
		// switch over the type of keyword.
		ch, _, err := r.Read()
		if err != nil {
			return rppb.Type_UNKNOWN, ""
		}

		if IsColon(ch) {
			err := r.Unread()
			if err != nil {
				fmt.Printf("failed to unread(%v) on reader.\n", string(ch))
				return rppb.Type_UNKNOWN, buf.String()
			}

			key := strings.ToUpper(buf.String())
			ret := buf.String()
			// Some RPSL keywords have '-' or '*' in them, the proto enum does not have
			// these characters included, remove them from the key comparison.
			key = strings.Replace(key, "-", "", -1)
			key = strings.Replace(key, "*", "", -1)

			if _, ok := rppb.Type_value[key]; !ok {
				return rppb.Type_UNKNOWN, ret
			}
			return rppb.Type(rppb.Type_value[key]), ret
		}
		// Each loop through accumulates more in the buffer.
		_, _ = buf.WriteRune(ch)
	}
}

// readValue reads the key's value, everything after the ":" and to a newline.
// Note that values can wrap to a second line:
//   if the newline is followed by a letter character the value is done.
//   if the newline is followed by a newline that is the end of the record.
//   if the newline is followed by a whitespace character the value continues.
// The bool in the return signals 'end of record'.
func (r *Reader) readValue() (string, bool, error) {
	var buf bytes.Buffer
	for {
		ch, _, err := r.Read()
		if err != nil {
			_, _ = buf.WriteRune(ch)
			return buf.String(), false, err
		}

		switch {
		// newline and letter, return value.
		case IsNewline(ch) && IsLetter(r.Peek()):
			return buf.String(), false, nil
			// newline and newline, return value and end-of-record.
		case IsNewline(ch) && (IsNewline(r.Peek()) || r.Peek() == eof):
			// read and discard the current newline and next newline (EOR).
			_, _, err := r.Read()
			if err != nil {
				return buf.String(), true, err
			}
			return buf.String(), true, nil
		}
		// nom-nom, keep on accumulating.
		_, _ = buf.WriteRune(ch)
	}
}

// Peek looks one character ahead, and return to the starting position.
func (r *Reader) Peek() rune {
	ch, _, err := r.Read()
	// Any errror during Peek, which is reading an already open file, is 'EOF', return eof.
	if err != nil {
		// Log the actual error for later.
		return eof
	}
	err = r.Unread()
	if err != nil {
		// Log the actual error for later.
		return eof
	}
	return ch
}

// addKV adds a key/value to the Fields field in a record.
func addKV(r *rppb.Record, k rppb.Type, v string) {
	r.Fields = append(r.Fields, &rppb.KeyValue{Key: k, Value: v})
}

// initRecord initializes the Record struct and gets the stores key/value.
func (r *Reader) initRecord() (*rppb.Record, error) {
	rec := &rppb.Record{Fields: []*rppb.KeyValue{}}

	key, literal := r.findKey()
	if key == rppb.Type_UNKNOWN {
		return nil, fmt.Errorf("failed to read a keyword found unexpected: %v\n", literal)
	}

	err := r.consumeColon()
	if err != nil {
		return rec, err
	}

	rec.Type = key
	val, re, err := r.readValue()
	if err != nil {
		if err == io.EOF {
			return nil, err
		}
		return nil, fmt.Errorf("failure reading value: %v\n", err)
	}

	// Add the key/value to the record as well.
	addKV(rec, key, val)
	if re {
		return nil, fmt.Errorf("Finished reading a record\n")
	}

	return rec, nil
}

// Parse parses through the content sending resulting records into a channel to the caller.
func Parse(rdr *Reader, rc chan<- *rppb.Record) {
	// Read the file content, return all accumulated records.
	// Return in case of parsing errors on record/keyword type.
	// Return in case of reading error.
	for {
		rec, err := rdr.initRecord()
		if err != nil {
			glog.Infof("failed to init the first Record: %v", err)
			return
		}

		for {
			key, literal := rdr.findKey()
			if key == rppb.Type_UNKNOWN {
				glog.Infof("failed to read a keyword found unexpected: %v\n", literal)
				break
			}

			err := rdr.consumeColon()
			if err != nil {
				rc <- rec
				glog.Infof("failed to consume a key's colon separator: %v", err)
				break
			}

			val, re, err := rdr.readValue()
			if err != nil {
				addKV(rec, key, val)
				rc <- rec
				if err == io.EOF {
					// EOF in a read means moving to the next file.
					glog.Infof("found an EOF while reading a value: %v", err)
					return
				}
				glog.Infof("failed to read a value: %v", err)
				return
			}

			// Add the key/value to the record as well.
			addKV(rec, key, val)

			if re {
				break
			}
		}
		rc <- rec
	}
}
