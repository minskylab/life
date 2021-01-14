package life

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/vektah/gqlparser/v2/ast"
)

const basics = `
directive @from(ref: String!) on FIELD_DEFINITION

directive @unique on FIELD_DEFINITION
directive @immutable on FIELD_DEFINITION
directive @nillable on FIELD_DEFINITION
directive @sensitive on FIELD_DEFINITION

directive @storageKey(key: String!) on FIELD_DEFINITION
directive @default(value: String!) on FIELD_DEFINITION
directive @updateDefault(value: String!) on FIELD_DEFINITION

directive @ent(name: String) on OBJECT

directive @out on OBJECT
`

func lifeSchemaSource() *ast.Source {
	return &ast.Source{
		Input:   basics,
		Name:    "goent.graphql",
		BuiltIn: true,
	}
}

func openSchemaSources(path string, withGoEntBasics bool) ([]*ast.Source, error) {
	fmt.Printf("path: %s\n", path)

	sources := []*ast.Source{}

	if withGoEntBasics {
		sources = append(sources, lifeSchemaSource())
	}

	files, err := filepath.Glob(path)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		data, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		sources = append(sources, &ast.Source{
			Input:   string(data),
			Name:    f,
			BuiltIn: false,
		})
	}

	return sources, nil
}
