package life

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/vektah/gqlparser/v2/ast"
)

const entBasePath = "entgo.io/ent"

func generateType(def *ast.Definition, enums map[string]*ast.Definition) *jen.File {
	m := jen.NewFile("schema")

	m.ImportName(entBasePath, "ent")
	m.ImportName(entBasePath+"/schema/field", "field")
	m.ImportName(entBasePath+"/schema/edge", "edge")
	// m.ImportName("github.com/minskylab/life/scalars", "scalars")

	m.Line()

	m.Comment(fmt.Sprintf("%s  holds the schema definition for the %s entity.", def.Name, def.Name))
	m.Type().Id(def.Name).Struct(
		jen.Qual(entBasePath, "Schema"),
	)

	forFields := []*ast.FieldDefinition{}
	forEdges := []*ast.FieldDefinition{}

	fieldEnumType := map[string][]string{}

	for _, field := range def.Fields {
		_, isKnownScalar := entScalars[field.Type.Name()]
		enumField, enumExists := enums[field.Type.Name()]
		if isKnownScalar {
			forFields = append(forFields, field)
		} else if enumExists {
			values := []string{}
			for _, val := range enumField.EnumValues {
				values = append(values, val.Name)
			}

			fieldEnumType[field.Name] = values

			forFields = append(forFields, field)
		} else {
			forEdges = append(forEdges, field)
		}
	}

	// lifeID := jen.Qual("github.com/minskylab/life/scalars", "ID").Call(jen.Lit(""))

	m.Comment(fmt.Sprintf("Fields of the %s.", def.Name))
	m.Func().Parens(jen.Id(def.Name)).Id("Fields").Params().
		Index().Qual(entBasePath, "Field").Block(
		jen.Return(
			jen.Index().Qual(entBasePath, "Field").ValuesFunc(func(g *jen.Group) {
				for _, field := range forFields {
					var fField *jen.Statement = jen.Line()

					fieldTypeName := field.Type.Name()
					enumValues := []string{}

					if values, exist := fieldEnumType[field.Name]; exist {
						fieldTypeName = "Enum"
						enumValues = values
					}

					fieldScalar := *entScalars[fieldTypeName]

					if fieldTypeName == "ID" { // special case ID scalar
						fField.Add(
							fieldScalar.Call(jen.Lit(field.Name)).
								Dot("NotEmpty").Call().
								Dot("Unique").Call().
								Dot("Immutable").Call(),
						)
						g.Add(fField)
						continue
					}

					if fieldTypeName == "Map" {
						fField.Add(fieldScalar.Call(jen.Lit(field.Name), jen.Map(jen.String()).Interface().Values()))
					} else {
						if field.Type.Elem != nil {
							if _, exist := entScalarArrays[fieldTypeName]; exist {
								fieldArrScalar := *entScalarArrays[fieldTypeName]
								fField.Add(fieldArrScalar.Call(jen.Lit(field.Name)))
							} else {
								fField.Add(fieldScalar.Call(jen.Lit(field.Name)))
							}
						} else {
							fField.Add(fieldScalar.Call(jen.Lit(field.Name)))
						}
					}

					if len(enumValues) > 0 {
						fField.Dot("Values").CallFunc(func(g *jen.Group) {
							for _, v := range enumValues {
								g.Add(jen.Lit(v))
							}
						})
					}

					if !field.Type.NonNull {
						fField.Dot("Optional").Call()
					}

					for _, directive := range field.Directives {
						applyFieldDirective(directive, fField)
					}

					g.Add(fField)
				}
				g.Line()
			}),
		),
	)

	m.Line()

	m.Comment(fmt.Sprintf("Edges of the %s.", def.Name))
	m.Func().Parens(jen.Id(def.Name)).Id("Edges").Params().
		Index().Qual(entBasePath, "Edge").Block(
		jen.Return(
			jen.Index().Qual(entBasePath, "Edge").ValuesFunc(func(g *jen.Group) {
				for _, field := range forEdges {
					var fEdge *jen.Statement = jen.Line()

					if field.Directives.ForName("from") != nil {
						fEdge.Qual(entBasePath+"/schema/edge", "From")
					} else if field.Directives.ForName("to") != nil {
						fEdge.Qual(entBasePath+"/schema/edge", "To")
					} else { // default
						fEdge.Qual(entBasePath+"/schema/edge", "To")
					}

					fEdge.Call(
						jen.Lit(field.Name),
						jen.Id(field.Type.Name()).Dot("Type"),
					)

					if field.Type.Elem == nil { // unique
						fEdge.Dot("Unique").Call()
					}

					if field.Type.NonNull {
						fEdge.Dot("Required").Call()
					}

					for _, directive := range field.Directives {
						applyEdgeDirective(directive, fEdge)
					}

					g.Add(fEdge)
				}
				g.Line()
			}),
		),
	)

	return m
}
