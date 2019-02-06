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
)

// Reader is a struct to manage access to the irr database file.
type Reader struct {
	reader *bufio.Reader
}

// Record is a single IRR record.
type Record struct {
	Type   KeyWord
	Fields map[KeyWord]string // Fields is the whole record indexed by key.
}

type Records []*Record

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
func (r *Reader) findKey() (KeyWord, string) {
	var buf bytes.Buffer
	for {
		// ReadRule until a colon is found, return the buf and
		// switch over the type of keyword.
		ch, _, err := r.Read()
		if err != nil {
			return EOF, ""
		}

		if IsColon(ch) {
			err := r.Unread()
			if err != nil {
				fmt.Printf("failed to unread(%v) on reader.\n", string(ch))
				return ILLEGAL, buf.String()
			}

			switch strings.ToUpper(buf.String()) {
			case "ADDRESS":
				return ADDRESS, buf.String()
			case "ADMIN-C":
				return ADMINC, buf.String()
			case "AGGR-BNDRY":
				return AGGRBNDRY, buf.String()
			case "AGGR-MTD":
				return AGGRMTD, buf.String()
			case "ALIAS":
				return ALIAS, buf.String()
			case "AS-NAME":
				return ASNAME, buf.String()
			case "AS-SET":
				return ASSET, buf.String()
			case "AUTH":
				return AUTH, buf.String()
			case "AUT-NUM":
				return AUTNUM, buf.String()
			case "CERTIF":
				return CERTIF, buf.String()
			case "CHANGED":
				return CHANGED, buf.String()
			case "COMPONENTS":
				return COMPONENTS, buf.String()
			case "COUNTRY":
				return COUNTRY, buf.String()
			case "DEFAULT":
				return DEFAULT, buf.String()
			case "DESCR":
				return DESCR, buf.String()
			case "E-MAIL":
				return EMAIL, buf.String()
			case "EOF":
				return EOF, buf.String()
			case "EXPORT":
				return EXPORT, buf.String()
			case "EXPORT-COMPS":
				return EXPORTCOMPS, buf.String()
			case "EXPORT-VIA":
				return EXPORTVIA, buf.String()
			case "FAX-NO":
				return FAXNO, buf.String()
			case "FILTER":
				return FILTER, buf.String()
			case "FILTER-SET":
				return FILTERSET, buf.String()
			case "FINGERPR":
				return FINGERPR, buf.String()
			case "GEOIDX":
				return GEOIDX, buf.String()
			case "HOLES":
				return HOLES, buf.String()
			case "IFADDR":
				return IFADDR, buf.String()
			case "IMPORT":
				return IMPORT, buf.String()
			case "IMPORT-VIA":
				return IMPORTVIA, buf.String()
			case "INET6NUM":
				return INET6NUM, buf.String()
			case "INETNUM":
				return INETNUM, buf.String()
			case "INET-RTR":
				return INETRTR, buf.String()
			case "INTERFACE":
				return INTERFACE, buf.String()
			case "KEY-CERT":
				return KEYCERT, buf.String()
			case "LOCAL-AS":
				return LOCALAS, buf.String()
			case "MBRS-BY-REF":
				return MBRSBYREF, buf.String()
			case "MEMBER-OF":
				return MEMBEROF, buf.String()
			case "MEMBERS":
				return MEMBERS, buf.String()
			case "METHOD":
				return METHOD, buf.String()
			case "MNT-BY":
				return MNTBY, buf.String()
			case "MNTNER":
				return MNTNER, buf.String()
			case "MNT-NFY":
				return MNTNFY, buf.String()
			case "MP-EXPORT":
				return MPEXPORT, buf.String()
			case "MP-FILTER":
				return MPFILTER, buf.String()
			case "MP-IMPORT":
				return MPIMPORT, buf.String()
			case "MP-MEMBERS":
				return MPMEMBERS, buf.String()
			case "MP-PEER":
				return MPPEER, buf.String()
			case "MP-PEERING":
				return MPPEERING, buf.String()
			case "NETNAME":
				return NETNAME, buf.String()
			case "NIC-HDL":
				return NICHDL, buf.String()
			case "NOTIFY":
				return NOTIFY, buf.String()
			case "ORIGIN":
				return ORIGIN, buf.String()
			case "OWNER":
				return OWNER, buf.String()
			case "PEER":
				return PEER, buf.String()
			case "PEERING":
				return PEERING, buf.String()
			case "PEERING-SET":
				return PEERINGSET, buf.String()
			case "PERSON":
				return PERSON, buf.String()
			case "PHONE":
				return PHONE, buf.String()
			case "REMARKS":
				return REMARKS, buf.String()
			case "ROA-URI":
				return ROAURI, buf.String()
			case "ROLE":
				return ROLE, buf.String()
			case "ROUTE":
				return ROUTE, buf.String()
			case "ROUTE6":
				return ROUTE6, buf.String()
			case "ROUTE-SET":
				return ROUTESET, buf.String()
			case "RS-IN":
				return RSIN, buf.String()
			case "RS-OUT":
				return RSOUT, buf.String()
			case "RTR-SET":
				return RTRSET, buf.String()
			case "SOURCE":
				return SOURCE, buf.String()
			case "STATUS":
				return STATUS, buf.String()
			case "TECH-C":
				return TECHC, buf.String()
			case "TROUBLE":
				return TROUBLE, buf.String()
			case "UPD-TO":
				return UPDTO, buf.String()
			case "*XXE":
				return XXE, buf.String()
			case "*XXNER":
				return XXNER, buf.String()
			case "*XX-NUM":
				return XXNUM, buf.String()
			case "*XXRING-SET":
				return XXRINGSET, buf.String()
			case "*XXSET":
				return XXSET, buf.String()
			case "*XXSON":
				return XXSON, buf.String()
			case "*XXTE":
				return XXTE, buf.String()
			case "*XXTE6":
				return XXTE6, buf.String()
			case "*XXTE-SET":
				return XXTESET, buf.String()
			default:
				// TODO(morrowc): Log this output instead of printing it to the console.
				// fmt.Printf("failed to match whatever is in buf currently: %v\n", buf.String())
				return ILLEGAL, buf.String()
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
		case IsNewline(ch) && IsLetter(r.peek()):
			return buf.String(), false, nil
			// newline and newline, return value and end-of-record.
		case IsNewline(ch) && (IsNewline(r.peek()) || r.peek() == eof):
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

// Look one character ahead, and return to the starting position.
func (r *Reader) peek() rune {
	ch, _, err := r.Read()
	// Any errror during peek, which is reading an already open file, is 'EOF', return eof.
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

// initRecord initializes the Record struct and gets the stores key/value.
func (r *Reader) initRecord() (*Record, error) {
	fieldMap := make(map[KeyWord]string)
	rec := &Record{Fields: fieldMap}

	key, literal := r.findKey()
	if key == ILLEGAL {
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
	rec.Fields[key] = val
	if re {
		return nil, fmt.Errorf("Finished reading a record\n")
	}

	return rec, nil
}

// Parse parses through the content creating records.
// TODO(morrowc): Convert this from sending back a single
// set of Records to pipelining parsed *Record records
// over a channel for coalation by the larger operating process.
func Parse(rdr *Reader) (Records, error) {
	// Create the record to fill.
	var rs Records

	// Read the file content, return all accumulated records.
	// Return in case of parsing errors on record/keyword type.
	// Return in case of reading error.
	for {
		rec, err := rdr.initRecord()
		if err != nil {
			return rs, err
		}

		for {
			key, literal := rdr.findKey()
			if key == ILLEGAL {
				return rs, fmt.Errorf("failed to read a keyword found unexpected: %v\n", literal)
			}

			err := rdr.consumeColon()
			if err != nil {
				rs = append(rs, rec)
				return rs, err
			}

			val, re, err := rdr.readValue()
			if err != nil {
				rec.Fields[key] = val
				rs = append(rs, rec)
				if err == io.EOF {
					return rs, err
				}
				return rs, err
			}

			// Add the key/value to the record as well.
			if _, ok := rec.Fields[key]; !ok {
				rec.Fields[key] = val
			} else {
				rec.Fields[key] = fmt.Sprintf("%v %v", rec.Fields[key], val)
			}

			if re {
				break
			}
		}
		// TODO(morrowc): As with above each record created
		// should be sent into a channel to be collected by the caller.
		rs = append(rs, rec)
	}

	return rs, nil
}
