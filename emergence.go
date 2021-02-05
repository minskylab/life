package life

import (
	"io"

	"github.com/vektah/gqlparser/v2/ast"
)

// EmergentGenerator receive an schema and returns a map with pathfiles as a key and io.Readers as values.
type EmergentGenerator func(schema *ast.Schema) (map[string]io.Reader, error)

// EmergentEffect describe an emergent side effect to generate graphql models from seed.
type EmergentEffect interface {
	OutputLocation() string
	Directives() *ast.Source
	Generator(schema *ast.Schema) (map[string]io.Reader, error)
}
