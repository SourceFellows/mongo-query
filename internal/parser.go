package internal

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

type (
	MongoDBStruct struct {
		Name          string
		BsonTag       string
		Fields        []*Field
		NestedStructs []*MongoDBStruct
	}

	Field struct {
		Name      string
		BsonTag   string
		ArrayType bool
	}
)

func ParseFile(filePath string, explicitStructs string) ([]*MongoDBStruct, error) {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, filePath, nil, 0)
	if err != nil {
		return nil, err
	}

	var mongoDbStructs []*MongoDBStruct
	for _, decl := range file.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}

		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}

			if explicitStructs == "" || strings.Contains(explicitStructs, ts.Name.Name) {
				mongoDbStruct := parseStruct(ts.Name.Name, st, file.Scope)
				mongoDbStructs = append(mongoDbStructs, mongoDbStruct)
			}
		}
	}

	return mongoDbStructs, nil
}

func parseStruct(name string, structType *ast.StructType, scope *ast.Scope) *MongoDBStruct {
	parsedFields, nestedStructs := parseFields(structType.Fields.List, scope)
	fields := make([]*Field, 0)
	for _, field := range parsedFields {
		if field != nil {
			fields = append(fields, field)
		}
	}

	return &MongoDBStruct{
		Name:          name,
		Fields:        fields,
		NestedStructs: nestedStructs,
	}
}

func parseFields(fields []*ast.Field, scope *ast.Scope) ([]*Field, []*MongoDBStruct) {
	parsedFields := make([]*Field, 0)
	parsedNestedStructs := make([]*MongoDBStruct, 0)

	for _, field := range fields {
		f := &Field{
			Name: field.Names[0].Name,
		}

		f.BsonTag = parseFieldName(field)
		if se, ok := field.Type.(*ast.StarExpr); ok {
			// overwrite type in case of pointer
			field.Type = se.X
		}

		switch field.Type.(type) {
		case *ast.Ident:
			it := field.Type.(*ast.Ident)
			if _, ok2 := scope.Objects[it.Name]; ok2 {
				nestedFields, nestedStructs := parseFields(it.Obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List, scope)
				parsedNestedStructs = append(parsedNestedStructs, &MongoDBStruct{
					Name:          f.Name,
					BsonTag:       f.BsonTag,
					Fields:        nestedFields,
					NestedStructs: nestedStructs,
				})
				continue
			}
			parsedFields = append(parsedFields, f)
		case *ast.StructType:
			nestedField, nestedStructs := parseFields(field.Type.(*ast.StructType).Fields.List, scope)
			parsedNestedStructs = append(parsedNestedStructs, &MongoDBStruct{
				Name:          f.Name,
				BsonTag:       f.BsonTag,
				Fields:        nestedField,
				NestedStructs: nestedStructs,
			})
		case *ast.ArrayType:
			it := field.Type.(*ast.ArrayType)
			f.ArrayType = true
			var nestedStructFields []*ast.Field
			switch it.Elt.(type) {
			case *ast.StructType:
				// inline struct definition
				nestedStructFields = it.Elt.(*ast.StructType).Fields.List
			case *ast.Ident:
				obj := it.Elt.(*ast.Ident).Obj
				// explicit struct definition
				if obj == nil {
					// add simple field in case of primitive array type
					parsedFields = append(parsedFields, f)
					continue
				}

				nestedStructFields = obj.Decl.(*ast.TypeSpec).Type.(*ast.StructType).Fields.List
			}

			nestedField, nestedStructs := parseFields(nestedStructFields, scope)
			parsedNestedStructs = append(parsedNestedStructs, &MongoDBStruct{
				Name:          f.Name,
				BsonTag:       f.BsonTag,
				Fields:        nestedField,
				NestedStructs: nestedStructs,
			})
			continue
		default:
			parsedFields = append(parsedFields, f)
		}

	}

	return parsedFields, parsedNestedStructs
}

func parseFieldName(field *ast.Field) string {
	fieldName := strings.ToLower(field.Names[0].Name)
	if field.Tag != nil {
		tag := field.Tag.Value
		if strings.Contains(tag, "bson") {
			splittedTags := strings.Split(tag, "\"")
			for j, splittedTag := range splittedTags {
				if strings.Contains(splittedTag, "bson") {
					fieldName = splittedTags[j+1]
					break
				}
			}
		}
	}

	return fieldName
}
