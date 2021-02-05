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

type dependenciesTree map[string][]string

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

func (effect *EmergentEffect) mutationsGenerator(schema *ast.Schema, tpl template.Template) (map[string]io.Reader, error) {
	files := map[string]io.Reader{}
	deps := dependenciesTree{}

	flatTypes := map[string]*ast.Definition{}

	for _, t := range schema.Types {
		flatTypes[t.Name] = t

		if t.Kind != ast.Object || strings.HasPrefix(t.Name, "__") {
			continue
		}

		if deps[t.Name] == nil {
			deps[t.Name] = []string{}
		}

		for _, f := range t.Fields {
			if isScalar(f.Type.Name()) {
				continue
			}

			deps[t.Name] = append(deps[t.Name], f.Type.Name())
		}

	}

	toGenerate := map[string]bool{}

	for typeName, typeDef := range flatTypes {
		if !isValidType(typeDef) {
			continue
		}

		toGenerate[typeDef.Name] = true

		for _, d := range deps[typeName] {
			dep, exist := flatTypes[d]
			if !exist {
				continue
			}

			toGenerate[dep.Name] = true
		}
	}

	for _, t := range schema.Types {
		if _, isValid := toGenerate[t.Name]; !isValid {
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
