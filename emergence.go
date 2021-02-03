package life

import (
	"html/template"
	"io"

	"github.com/vektah/gqlparser/v2/ast"
)

// EmergentGenerator receive an schema and returns a map with pathfiles as a key and io.Readers as values.
type EmergentGenerator func(schema *ast.Schema, tpl template.Template) (map[string]io.Reader, error)

// EmergentEffect describe an emergent side effect to generate graphql models from seed.
type EmergentEffect struct {
	OutputLocation string
	Directives     *ast.Source
	Generator      EmergentGenerator
}
