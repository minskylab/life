package life

import (
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func generate(filepath string, folderOut string, withGoEntBasics bool) error {
	sources, err := openSchemaSources(filepath, withGoEntBasics)
	if err != nil {
		return errors.WithStack(err)
	}

	sch, gqlErr := gqlparser.LoadSchema(sources...)
	if gqlErr != nil {
		return errors.WithStack(gqlErr)
	}

	defs := []*ast.Definition{}
	enums := map[string]*ast.Definition{}

	for _, t := range sch.Types {
		if t.Name == "Query" { // make something with mutations and queries
			continue
		} else if t.Name == "Mutation" {
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
			if t.Directives.ForName("response") != nil {
				continue
			}

			if t.Directives.ForName("out") != nil {
				continue
			}

			if t.Directives.ForName("exclude") != nil {
				continue
			}

			defs = append(defs, t)
		}
	}

	return generateTypes(folderOut, defs, enums)
}
