package life

import (
	"io"

	"github.com/vektah/gqlparser/v2/ast"
)

// EmergentGenerator receive an schema and returns a map with pathfiles as a key and io.Readers as values.
type EmergentGenerator func(schema *ast.Schema) map[string]io.Reader

// EmergentEffect describe an emergent side effect to generate graphql models from seed.
type EmergentEffect struct {
	RelativeLocation string
	Directives       *ast.Source
	Generator        EmergentGenerator
}
