/**
 * MIT License
 *
 * Copyright (c) 2023 Source Fellows GmbH (https://www.source-fellows.com)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package internal

import (
	"bytes"
	_ "embed"
	"strings"
	"testing"
)

//go:embed example-types_test.go
var testFile []byte

func TestParseFile(t *testing.T) {
	//given
	r := bytes.NewReader(testFile)

	//when
	mongoDBStructs, err := ParseFile(r, "")

	//then
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	if len(mongoDBStructs) != 4 {
		t.Errorf("expected to have 4 struct but had %d", len(mongoDBStructs))
	}

}

func TestParseFile_withExplicitStructs(t *testing.T) {
	//given
	r := bytes.NewReader(testFile)
	structs := "Basic,WithArray"

	//when
	mongoDBStructs, err := ParseFile(r, structs)

	//then
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	if len(mongoDBStructs) != 2 {
		t.Errorf("expected to have 2 struct but had %d", len(mongoDBStructs))
		return
	}

	for _, dbStruct := range mongoDBStructs {
		if !strings.Contains(structs, dbStruct.Name) {
			t.Errorf("got unexpected mongodb struct %s", dbStruct.Name)
			return
		}
	}
}
