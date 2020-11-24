package codegen

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/vektah/gqlparser/v2/ast"
)

func generateTypes(definitions []*ast.Definition) {
	m := jen.NewFile("models")

	for _, def := range definitions {
		fields := []jen.Code{}
		for _, field := range def.Fields {
			name := strings.ToTitle(field.Name)
			fields = append(fields, jen.Id(name).String())
		}

		m.Type().Id(def.Name).Struct(fields...).Line()
	}

	fmt.Printf("%#v", m)
}
