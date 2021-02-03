package automutations

import (
	"fmt"
	"html/template"

	"github.com/minskylab/life"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/ast"
)

// EmergentEffect wraps a life emergent effect.
type EmergentEffect struct {
	life.EmergentEffect
	tpl *template.Template
}

// NewAutoMutationEffect ...
func NewAutoMutationEffect(location string) (*EmergentEffect, error) {
	// wd, err := os.Getwd()
	// if err != nil {
	// 	return nil, errors.WithStack(err)
	// }

	// templateFilepath := path.Join(wd, "effects", "automutations", "extention.graphql.tpl")

	// ioutil.

	tpl, err := template.New("automutation").Parse(templateString)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &EmergentEffect{
		tpl: tpl,
		EmergentEffect: life.EmergentEffect{
			OutputLocation: location,
			Directives: &ast.Source{
				Name:    "automutations.graphql",
				Input:   fmt.Sprintf("directive @%s on FIELD_DEFINITION", exposedDirectiveName),
				BuiltIn: true,
			},
			Generator: mutationsGenerator,
		},
	}, nil
}
