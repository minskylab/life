package life

import (
	"io/ioutil"
	"path/filepath"

	"github.com/vektah/gqlparser/v2/ast"
)

const basics = `
directive @unique on FIELD_DEFINITION
directive @from(ref: String!) on FIELD_DEFINITION
directive @immutable on FIELD_DEFINITION
directive @default(value: String!) on FIELD_DEFINITION
directive @response on OBJECT
directive @out on OBJECT
directive @exclude on OBJECT
`

func lifeSchemaSource() *ast.Source {
	return &ast.Source{
		Input:   basics,
		Name:    "goent.graphql",
		BuiltIn: true,
	}
}

func openSchemaSources(path string) ([]*ast.Source, error) {
	// info, err := os.Stat(path)
	// if err != nil {
	// 	return nil, err
	// }

	// if os.IsNotExist(err) {
	// 	return nil, errors.New("file/dir does not exist")
	// }

	sources := []*ast.Source{
		lifeSchemaSource(),
	}

	// if info.IsDir() {
	// 	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

	// 	})
	// }

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
