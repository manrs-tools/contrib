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

package rpsl

// IsWhitespace validates if the rune is a space or tab.
func IsWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// IsNewline validates if the rune is a newline.
func IsNewline(ch rune) bool { return ch == '\n' }

// IsColon validates that the rune is a ':'.
func IsColon(ch rune) bool { return (ch == colon) }

// IsOctothorpe validates that the current rune is a '#'.
func IsOctothorpe(ch rune) bool { return (ch == octothorpe) }

// ConsumeComment reads a full octothorpe intiated comment, and discards it.
func (r *Reader) ConsumeComment() error {
	if IsOctothorpe(r.Peek()) {
		_, _, err := r.Read()
		if err != nil {
			return err
		}

		err = r.consumeToNewline()
		if err != nil {
			return err
		}

		// If the next rune is a '#', read the next comment as well
		err = r.ConsumeComment()
		if err != nil {
			return err
		}
	}
	return nil
}

// Consume leading whitespace from the reader.
func (r *Reader) ConsumeLeadingWS() error {
	for {
		if !IsWhitespace(r.Peek()) {
			return nil
		}
		_, _, err := r.Read()
		if err != nil {
			return err
		}
	}
}

// consumeColon reads the colon after a key and all whitespace before the value.
func (r *Reader) consumeColon() error {
	if IsColon(r.Peek()) {
		_, _, err := r.Read()
		if err != nil {
			return err
		}

		err = r.ConsumeLeadingWS()
		if err != nil {
			return err
		}
	}
	return nil
}

// consumeToNewline reads all content until a newline rune is encountered.
func (r *Reader) consumeToNewline() error {
	for {
		if IsNewline(r.Peek()) {
			_, _, err := r.Read()
			if err != nil {
				return err
			}
			return nil
		}
		_, _, err := r.Read()
		if err != nil {
			return err
		}
	}
}
