package automutations

import (
	"fmt"
	"io"

	"github.com/minskylab/life"
	"github.com/vektah/gqlparser/v2/ast"
)

const exposedDirectiveName = "exposed"

func mutationsGenerator(schema *ast.Schema) map[string]io.Reader {
	files := map[string]io.Reader{}

	for _, t := range schema.Types {
		if t.Kind == ast.Object {
			entity := generateStructures(t)

			fmt.Println(entity)
		}
	}

	return files
}

// NewAutoMutation ...
func NewAutoMutation(location string) life.EmergentEffect {
	return life.EmergentEffect{
		RelativeLocation: location,
		Directives: &ast.Source{
			Name:    "automutations.graphql",
			Input:   fmt.Sprintf("directive @%s on FIELD_DEFINITION", exposedDirectiveName),
			BuiltIn: true,
		},
		Generator: mutationsGenerator,
	}
}
