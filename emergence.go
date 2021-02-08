package life

import (
	"io"

	"github.com/vektah/gqlparser/v2/ast"
)

// EmergentEffect describe an emergent side effect to generate graphql models from seed.
type EmergentEffect interface {
	OutputLocation() string
	Directives() *ast.Source
	Generator(schema *ast.Schema) (map[string]io.Reader, error)
}
