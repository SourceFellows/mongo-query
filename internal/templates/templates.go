package templates

import _ "embed"

//go:embed mongoFilter-template.gotmpl
var MongoFilterTemplate string

//go:embed structFilter-template.gotmpl
var StructFilterTemplate string
