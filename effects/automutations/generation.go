package automutations

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/ast"
)

type void struct{}

var member void

type dependencies map[string]void

const exposedDirectiveName = "exposed"

func isValidType(t *ast.Definition) bool {
	if t.Directives.ForName(exposedDirectiveName) == nil {
		return false
	}

	if t.Kind != ast.Object {
		return false
	}

	if strings.HasPrefix(t.Name, "_") || t.Name == "Query" || t.Name == "Mutation" || t.Name == "Subscription" {
		return false
	}

	return true
}

func (effect *EmergentEffect) recursiveDeps(current *ast.Definition, types map[string]*ast.Definition, deps dependencies) dependencies {
	if deps == nil {
		deps = dependencies{}
	}

	if current == nil {
		for tName, t := range types {
			if isValidType(t) {
				deps[tName] = member

				newDeps := effect.recursiveDeps(t, types, deps)
				for d := range newDeps {
					deps[d] = member
				}
			}
		}
	} else {
		for _, f := range current.Fields {
			fName := f.Type.Name()
			if isScalar(fName) {
				continue
			}

			_, exists := deps[fName]

			if exists || types[fName].Kind != ast.Object {
				continue
			}

			deps[fName] = member

			newDeps := effect.recursiveDeps(types[fName], types, deps)
			for d := range newDeps {
				deps[d] = member
			}

			// return deps
		}
	}

	return deps
}

func (effect *EmergentEffect) mutationsGenerator(schema *ast.Schema, tpl template.Template) (map[string]io.Reader, error) {
	files := map[string]io.Reader{}
	// deps := dependenciesTree{}

	flatTypes := map[string]*ast.Definition{}

	for _, t := range schema.Types {
		flatTypes[t.Name] = t
	}

	deps := effect.recursiveDeps(nil, flatTypes, nil)

	for _, t := range schema.Types {
		if _, isValid := deps[t.Name]; !isValid {
			continue
		}

		entity := generateStructure(t)

		buff := bytes.NewBufferString("")

		if err := tpl.Execute(buff, entity); err != nil {
			return nil, errors.WithStack(err)
		}

		files[fmt.Sprintf("%s.graphql", strings.ToLower(t.Name))] = buff
	}

	return files, nil
}
