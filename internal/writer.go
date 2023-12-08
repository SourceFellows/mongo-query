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
