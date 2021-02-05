package life

import (
	"io"
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func generate(filepath string, folderOut string, opts GenerationOptions) error {
	sources, err := openSchemaSources(filepath, opts.EntDirectivesBuiltIn)
	if err != nil {
		return errors.WithStack(err)
	}

	for _, effect := range opts.Effects {
		sources = append(sources, effect.Directives())
	}

	sch, gqlErr := gqlparser.LoadSchema(sources...)
	if gqlErr != nil {
		return errors.WithStack(gqlErr)
	}

	defs := []*ast.Definition{}
	enums := map[string]*ast.Definition{}

	for _, t := range sch.Types {
		// make something (someday) with mutations and queries
		if t.Name == "Query" {
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
			// collect enums
			enums[t.Name] = t
		} else if t.Kind == ast.Interface {
			continue
		} else if t.Kind == ast.Scalar {
			continue
		} else if t.Kind == ast.Union {
			continue
		} else {
			if t.Directives.ForName("out") != nil {
				continue
			}

			if t.Directives.ForName("ent") == nil {
				continue
			}

			defs = append(defs, t)
		}
	}

	if err := generateTypes(folderOut, defs, enums); err != nil {
		return errors.WithStack(err)
	}

	for _, effect := range opts.Effects {
		// effect.Generator()()
		generativeEffect, err := effect.Generator(sch)
		if err != nil {
			return errors.WithStack(err)
		}

		for fName, emergence := range generativeEffect {
			filepath := path.Join(effect.OutputLocation(), fName)

			// fmt.Println("emergence: " + filepath)

			file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR, 0644)
			if err != nil {
				return errors.WithStack(err)
			}

			_ = file.Truncate(0) // delete all current content

			if _, err = io.Copy(file, emergence); err != nil {
				return errors.WithStack(err)
			}

			if err := file.Close(); err != nil {
				return errors.WithStack(err)
			}
		}
	}

	return nil
}
