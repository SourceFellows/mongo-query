package internal

import (
	"bytes"
	"github.com/sourcefellows/mongo-query/internal/templates"
	"go/format"
	"html/template"
	"io"
	"strings"
)

type WriterType struct {
	template string
}

var (
	StructWriter = WriterType{templates.StructFilterTemplate}
	StaticWriter = WriterType{templates.MongoFilterTemplate}
)

func Write(dbStruct *MongoDBStruct, writerType WriterType, writer io.Writer) error {
	funcs := template.FuncMap{
		"toLower": toLower,
	}

	tmpl, err := template.New("type").
		Funcs(funcs).
		Parse(writerType.template)
	if err != nil {
		return err
	}

	var bites bytes.Buffer
	err = tmpl.Execute(&bites, dbStruct)
	if err != nil {
		return err
	}

	src, err := format.Source(bites.Bytes())
	if err != nil {
		return err
	}

	_, err = writer.Write(src)

	return err
}

func toLower(input string) string {
	return strings.ToLower(input)
}
