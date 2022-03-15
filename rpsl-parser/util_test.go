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

// Tests for the util.go utilities.
package rpsl

import (
	"strings"
	"testing"
)

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Sucess WS ' '",
		char: rune(' '),
		want: true,
	}, {
		desc: "Success WS '\t'",
		char: rune('\t'),
		want: true,
	}, {
		desc: "Fail WS a",
		char: rune('a'),
		want: false,
	}}

	for _, test := range tests {
		got := IsWhitespace(test.char)
		if test.want != got {
			t.Errorf("[%v]: failed got: %v want: %v", test.desc, got, test.want)
		}
	}
}

func TestIsNewline(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Sucess NL '\n'",
		char: rune('\n'),
		want: true,
	}, {
		desc: "Fail NL '\t'",
		char: rune('\t'),
		want: false,
	}, {
		desc: "Fail NL a",
		char: rune('a'),
		want: false,
	}}

	for _, test := range tests {
		got := IsNewline(test.char)
		if test.want != got {
			t.Errorf("[%v]: failed got: %v want: %v", test.desc, got, test.want)
		}
	}
}

func TestIsColon(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Sucess Colon :",
		char: rune(':'),
		want: true,
	}, {
		desc: "Fail Colon '\t'",
		char: rune('\t'),
		want: false,
	}, {
		desc: "Fail Colon ;",
		char: rune(';'),
		want: false,
	}}

	for _, test := range tests {
		got := IsColon(test.char)
		if test.want != got {
			t.Errorf("[%v]: failed got: %v want: %v", test.desc, got, test.want)
		}
	}
}

func TestIsOctothorpe(t *testing.T) {
	tests := []struct {
		desc string
		char rune
		want bool
	}{{
		desc: "Sucess Octothorpe #",
		char: rune('#'),
		want: true,
	}, {
		desc: "Fail Octothorpe '\t'",
		char: rune('\t'),
		want: false,
	}, {
		desc: "Fail Octothorpe ;",
		char: rune(';'),
		want: false,
	}}

	for _, test := range tests {
		got := IsOctothorpe(test.char)
		if test.want != got {
			t.Errorf("[%v]: failed got: %v want: %v", test.desc, got, test.want)
		}
	}
}

func TestConsumeComment(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		want    rune
		wantErr bool
	}{{
		desc:  "Success",
		input: "# this is a comment\naut-num:",
		want:  rune('a'),
	}, {
		desc:  "Success 2 line comment",
		input: "# this is one line\n# this is the second line\naut-num:",
		want:  rune('a'),
	}, {
		desc:  "Success no comment",
		input: "aut-num:",
		want:  rune('a'),
	}, {
		desc:    "Fail EOF",
		input:   "# this is one line",
		wantErr: true,
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		err := r.ConsumeComment()
		got := r.Peek()
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		case err == nil:
			if got != test.want {
				t.Errorf("[%v]: got/want mismatched: %v/%v", test.desc, got, test.want)
			}
		}
	}
}

func TestConsumeLeadingWS(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		wantErr bool
	}{{
		desc:  "Successful lead WS read",
		input: "  foobar",
	}, {
		desc:  "Successful no lead WS read",
		input: "f oobar",
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		err := r.ConsumeLeadingWS()
		switch {
		case err != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, err)
		case err == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		}
	}
}

func TestConsumeColon(t *testing.T) {
	// Note that consumeColon will not error.
	// Providing a : and nothing else will still return no errors.
	tests := []struct {
		desc    string
		input   string
		wantErr bool
	}{{
		desc:  "Success",
		input: ":",
	}, {
		desc:  "Sucess colon and whitespace (space)",
		input: ": ",
	}, {
		desc: "Sucess colon and whitespace (tab)",
		input: ":	f",
	}, {
		desc:  "Success colon and no whitespace",
		input: ":f",
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		got := r.consumeColon()

		switch {
		case got != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, got)
		case got == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		}
	}
}

func TestConsumeToNewLine(t *testing.T) {
	tests := []struct {
		desc    string
		input   string
		wantErr bool
	}{{
		desc:  "Success",
		input: "foo bar\n",
	}, {
		desc:    "Fail no newline",
		input:   "foobar",
		wantErr: true,
	}}

	for _, test := range tests {
		r := NewReader(strings.NewReader(test.input))
		got := r.consumeToNewline()

		switch {
		case got != nil && !test.wantErr:
			t.Errorf("[%v]: got error when not expecting one: %v", test.desc, got)
		case got == nil && test.wantErr:
			t.Errorf("[%v]: did not get error, expected one.", test.desc)
		}
	}
}
