package life

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/vektah/gqlparser/v2/ast"
)

func openSchemaSources(path string, withGoEntBasics bool) ([]*ast.Source, error) {
	fmt.Printf("path: %s\n", path)

	sources := []*ast.Source{}

	if withGoEntBasics {
		sources = append(sources, lifeDirectivesSource())
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
