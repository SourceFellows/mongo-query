package internal

import (
	"strings"
	"testing"
)

func TestParseFile(t *testing.T) {
	//given
	filePath := "examples/types.go"

	//when
	mongoDBStructs, err := ParseFile(filePath, "")

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
	filePath := "examples/types.go"
	structs := "Basic,WithArray"

	//when
	mongoDBStructs, err := ParseFile(filePath, structs)

	//then
	if err != nil {
		t.Errorf("unexpected error %v", err)
		return
	}

	if len(mongoDBStructs) != 2 {
		t.Errorf("expected to have 2 struct but had %d", len(mongoDBStructs))
	}

	for _, dbStruct := range mongoDBStructs {
		if !strings.Contains(structs, dbStruct.Name) {
			t.Errorf("got unexpected mongodb struct %s", dbStruct.Name)
			return
		}
	}
}
