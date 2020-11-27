package life

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser"
	"github.com/vektah/gqlparser/ast"
)

func generate(filepath string, folderOut string) error {

	sources, err := openSchemaSources(filepath)
	if err != nil {
		return errors.WithStack(err)
	}

	sch, gqlErr := gqlparser.LoadSchema(sources...)
	if gqlErr != nil {
		return errors.WithStack(gqlErr)
	}

	defs := []*ast.Definition{}
	enums := map[string]*ast.Definition{}
	// queries := []*ast.Definition{}
	// mutations := []*ast.Definition{}

	for _, t := range sch.Types {
		if t.Name == "Query" { // make something with mutations and queries
			// queries = append(queries, t)
			continue
		} else if t.Name == "Mutation" {
			// mutations = append(mutations, t)
			continue
		} else if t.Name == "Subscription" {
			continue
		} else if t.BuiltIn {
			continue
		} else if t.Kind == ast.InputObject {
			continue
		} else if t.Kind == ast.Enum {
			enums[t.Name] = t
		} else if t.Kind == ast.Interface {
			continue
		} else if t.Kind == ast.Scalar {
			continue
		} else if t.Kind == ast.Union {
			continue
		} else {
			defs = append(defs, t)
		}
	}

	return generateTypes(folderOut, defs, enums)
}
