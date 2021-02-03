package automutations

import (
	"bytes"
	"html/template"
	"io"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/ast"
)

const exposedDirectiveName = "exposed"

func isValidType(t *ast.Definition) bool {
	if t.Directives.ForName(exposedDirectiveName) == nil {
		return false
	}

	if strings.HasPrefix(t.Name, "_") || t.Name == "Query" || t.Name == "Mutation" || t.Name == "Subscription" {
		return false
	}

	return true
}

func mutationsGenerator(schema *ast.Schema, tpl template.Template) (map[string]io.Reader, error) {
	files := map[string]io.Reader{}

	for _, t := range schema.Types {
		if t.Kind == ast.Object {
			if !isValidType(t) {
				continue
			}
			entity := generateStructures(t)
			buff := bytes.NewBufferString("")
			if err := tpl.Execute(buff, entity); err != nil {
				return nil, errors.WithStack(err)
			}
			spew.Dump(entity)
		}
	}

	return files, nil
}
