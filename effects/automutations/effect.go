package automutations

import (
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/ast"
)

// EmergentEffect wraps a life emergent effect.
type EmergentEffect struct {
	outputLocation string
	tpl            *template.Template
}

// OutputLocation implements a life emergent effect
func (effect *EmergentEffect) OutputLocation() string {
	return effect.outputLocation
}

// Directives implements a life emergent effect
func (effect *EmergentEffect) Directives() *ast.Source {
	return &ast.Source{
		Name:    "automutations.graphql",
		Input:   fmt.Sprintf("directive @%s on FIELD_DEFINITION", exposedDirectiveName),
		BuiltIn: true,
	}
}

// Generator implements a life emergent effect
func (effect *EmergentEffect) Generator(schema *ast.Schema) (map[string]io.Reader, error) {
	return effect.mutationsGenerator(schema, *effect.tpl)
}

// NewAutoMutationEffect ...
func NewAutoMutationEffect(location string) (*EmergentEffect, error) {
	tpl, err := template.New("automutation").Funcs(template.FuncMap{
		"title": func(text string) string {
			return strings.Title(text)
		},
	}).Parse(templateString)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &EmergentEffect{
		tpl:            tpl,
		outputLocation: location,
	}, nil
}
