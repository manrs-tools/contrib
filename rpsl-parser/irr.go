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
	"errors"
	"fmt"
	"io"
	"strings"
	"sync"
	"unicode"

	"github.com/golang/glog"
	rppb "github.com/manrs-tools/contrib/rpsl-parser/proto"
)

type Records struct {
	Records []*rppb.Record
	mu      sync.Mutex
}

func (r *Records) AddRecord(record *rppb.Record) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Records = append(r.Records, record)
}

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

			switch strings.ToUpper(buf.String()) {
			case "ROUTE":
				return rppb.Type_ROUTE, buf.String()
			case "ROUTE6":
				return rppb.Type_ROUTE6, buf.String()
			case "AUT-NUM":
				return rppb.Type_AUTNUM, buf.String()
			case "AS-SET":
				return rppb.Type_ASSET, buf.String()
			case "ROUTE-SET":
				return rppb.Type_ROUTESET, buf.String()
			case "MNTNER":
				return rppb.Type_MNTNER, buf.String()
			case "PERSON":
				return rppb.Type_PERSON, buf.String()
			case "INETNUM":
				return rppb.Type_INETNUM, buf.String()
			case "KEY-CERT":
				return rppb.Type_KEYCERT, buf.String()
			case "ROLE":
				return rppb.Type_ROLE, buf.String()
			case "INET-RTR":
				return rppb.Type_INETRTR, buf.String()
			case "INET6NUM":
				return rppb.Type_INET6NUM, buf.String()
			case "FILTER-SET":
				return rppb.Type_FILTERSET, buf.String()
			case "RTR-SET":
				return rppb.Type_RTRSET, buf.String()
			case "PEERING-SET":
				return rppb.Type_PEERINGSET, buf.String()
			case "ADDRESS":
				return rppb.Type_ADDRESS, buf.String()
			case "ADMIN-C":
				return rppb.Type_ADMINC, buf.String()
			case "AGGR-BNDRY":
				return rppb.Type_AGGRBNDRY, buf.String()
			case "AGGR-MTD":
				return rppb.Type_AGGRMTD, buf.String()
			case "ALIAS":
				return rppb.Type_ALIAS, buf.String()
			case "AS-NAME":
				return rppb.Type_ASNAME, buf.String()
			case "AUTH":
				return rppb.Type_AUTH, buf.String()
			case "CERTIF":
				return rppb.Type_CERTIF, buf.String()
			case "CHANGED":
				return rppb.Type_CHANGED, buf.String()
			case "COMPONENTS":
				return rppb.Type_COMPONENTS, buf.String()
			case "COUNTRY":
				return rppb.Type_COUNTRY, buf.String()
			case "DEFAULT":
				return rppb.Type_DEFAULT, buf.String()
			case "DESCR":
				return rppb.Type_DESCR, buf.String()
			case "E-MAIL":
				return rppb.Type_EMAIL, buf.String()
			case "EXPORT":
				return rppb.Type_EXPORT, buf.String()
			case "EXPORT-COMPS":
				return rppb.Type_EXPORTCOMPS, buf.String()
			case "EXPORT-VIA":
				return rppb.Type_EXPORTVIA, buf.String()
			case "FAX-NO":
				return rppb.Type_FAXNO, buf.String()
			case "FILTER":
				return rppb.Type_FILTER, buf.String()
			case "FINGERPR":
				return rppb.Type_FINGERPR, buf.String()
			case "GEOIDX":
				return rppb.Type_GEOIDX, buf.String()
			case "HOLES":
				return rppb.Type_HOLES, buf.String()
			case "IFADDR":
				return rppb.Type_IFADDR, buf.String()
			case "IMPORT":
				return rppb.Type_IMPORT, buf.String()
			case "IMPORT-VIA":
				return rppb.Type_IMPORTVIA, buf.String()
			case "INTERFACE":
				return rppb.Type_INTERFACE, buf.String()
			case "LOCAL-AS":
				return rppb.Type_LOCALAS, buf.String()
			case "MBRS-BY-REF":
				return rppb.Type_MBRSBYREF, buf.String()
			case "MEMBER-OF":
				return rppb.Type_MEMBEROF, buf.String()
			case "MEMBERS":
				return rppb.Type_MEMBERS, buf.String()
			case "METHOD":
				return rppb.Type_METHOD, buf.String()
			case "MNT-BY":
				return rppb.Type_MNTBY, buf.String()
			case "MNT-NFY":
				return rppb.Type_MNTNFY, buf.String()
			case "MP-EXPORT":
				return rppb.Type_MPEXPORT, buf.String()
			case "MP-FILTER":
				return rppb.Type_MPFILTER, buf.String()
			case "MP-IMPORT":
				return rppb.Type_MPIMPORT, buf.String()
			case "MP-MEMBERS":
				return rppb.Type_MPMEMBERS, buf.String()
			case "MP-PEER":
				return rppb.Type_MPPEER, buf.String()
			case "MP-PEERING":
				return rppb.Type_MPPEERING, buf.String()
			case "NETNAME":
				return rppb.Type_NETNAME, buf.String()
			case "NIC-HDL":
				return rppb.Type_NICHDL, buf.String()
			case "NOTIFY":
				return rppb.Type_NOTIFY, buf.String()
			case "ORIGIN":
				return rppb.Type_ORIGIN, buf.String()
			case "OWNER":
				return rppb.Type_OWNER, buf.String()
			case "PEER":
				return rppb.Type_PEER, buf.String()
			case "PEERING":
				return rppb.Type_PEERING, buf.String()
			case "PHONE":
				return rppb.Type_PHONE, buf.String()
			case "REMARKS":
				return rppb.Type_REMARKS, buf.String()
			case "ROA-URI":
				return rppb.Type_ROAURI, buf.String()
			case "RS-IN":
				return rppb.Type_RSIN, buf.String()
			case "RS-OUT":
				return rppb.Type_RSOUT, buf.String()
			case "SOURCE":
				return rppb.Type_SOURCE, buf.String()
			case "STATUS":
				return rppb.Type_STATUS, buf.String()
			case "TECH-C":
				return rppb.Type_TECHC, buf.String()
			case "TROUBLE":
				return rppb.Type_TROUBLE, buf.String()
			case "UPD-TO":
				return rppb.Type_UPDTO, buf.String()
			case "*XXE":
				return rppb.Type_XXE, buf.String()
			case "*XXNER":
				return rppb.Type_XXNER, buf.String()
			case "*XX-NUM":
				return rppb.Type_XXNUM, buf.String()
			case "*XXRING-SET":
				return rppb.Type_XXRINGSET, buf.String()
			case "*XXSET":
				return rppb.Type_XXSET, buf.String()
			case "*XXSON":
				return rppb.Type_XXSON, buf.String()
			case "*XXTE":
				return rppb.Type_XXTE, buf.String()
			case "*XXTE6":
				return rppb.Type_XXTE6, buf.String()
			case "*XXTE-SET":
				return rppb.Type_XXTESET, buf.String()
			default:
				// TODO(morrowc): Log this output instead of printing it to the console.
				// fmt.Printf("failed to match whatever is in buf currently: %v\n", buf.String())
				return rppb.Type_UNKNOWN, buf.String()
			}

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
		case IsNewline(ch) && unicode.IsLetter(r.Peek()):
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
		return nil, fmt.Errorf("failed to read a keyword found unexpected: %v", literal)
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
		return nil, fmt.Errorf("failure reading value: %v", err)
	}

	// Add the key/value to the record as well.
	addKV(rec, key, val)
	if re {
		return nil, errors.New("finished reading a record")
	}

	return rec, nil
}

// Parse parses through the content sending resulting records into a channel to the caller.
func Parse(rdr *Reader, records *Records) {
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
				records.AddRecord(rec)
				glog.Infof("failed to consume a key's colon separator: %v", err)
				break
			}

			val, re, err := rdr.readValue()
			if err != nil {
				addKV(rec, key, val)
				records.AddRecord(rec)
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
		records.AddRecord(rec)
	}
}
