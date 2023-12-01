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
