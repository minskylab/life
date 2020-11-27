package life

import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/vektah/gqlparser/ast"
)

func generateType(def *ast.Definition, enums map[string]*ast.Definition) *jen.File {
	m := jen.NewFile("schema")

	m.ImportName("github.com/facebook/ent", "ent")
	m.ImportName("github.com/facebook/ent/schema/field", "field")
	m.ImportName("github.com/facebook/ent/schema/edge", "edge")
	m.ImportName("github.com/minskylab/life/scalars", "scalars")

	m.Line()

	m.Comment(fmt.Sprintf("%s  holds the schema definition for the %s entity.", def.Name, def.Name))
	m.Type().Id(def.Name).Struct(
		jen.Qual("github.com/facebook/ent", "Schema"),
	)

	forFields := []*ast.FieldDefinition{}
	forEdges := []*ast.FieldDefinition{}

	for _, field := range def.Fields {
		_, isKnownScalar := entScalars[field.Type.Name()]
		_, enumExists := enums[field.Type.Name()]
		if isKnownScalar {
			forFields = append(forFields, field)
		} else if enumExists {
			// enums
		} else {
			forEdges = append(forEdges, field)
		}
	}

	// lifeID := jen.Qual("github.com/minskylab/life/scalars", "ID").Call(jen.Lit(""))

	m.Comment(fmt.Sprintf("Fields of the %s.", def.Name))
	m.Func().Parens(jen.Id(def.Name)).Id("Fields").Params().
		Index().Qual("github.com/facebook/ent", "Field").Block(
		jen.Return(
			jen.Index().Qual("github.com/facebook/ent", "Field").ValuesFunc(func(g *jen.Group) {
				for _, field := range forFields {
					var f *jen.Statement = jen.Line()
					fieldTypeName := field.Type.Name()

					fieldScalar := *entScalars[fieldTypeName]

					if fieldTypeName == "ID" { // special case ID scalar
						f.Add(
							fieldScalar.Call(jen.Lit(field.Name)).
								Dot("NotEmpty").Call().
								Dot("Unique").Call().
								Dot("Immutable").Call(),
						)
						g.Add(f)
						continue
					}

					if fieldTypeName == "Map" {
						f.Add(fieldScalar.Call(jen.Lit(field.Name), jen.Map(jen.String()).Interface().Values()))
						// } else if fieldTypeName == "ID" {
						// 	f.Add(fieldScalar.Call(jen.Lit(field.Name), lifeID))
					} else {
						f.Add(fieldScalar.Call(jen.Lit(field.Name)))
					}

					if !field.Type.NonNull {
						f.Dot("Optional").Call()
					}

					// f.Render(os.Stdout)
					// if field.Type

					for _, directive := range field.Directives {
						switch directive.Name {
						case "unique":
							f.Dot("Unique").Call()
						case "immutable":
							f.Dot("Immutable").Call()
						case "default":
							valueArg := directive.Arguments.ForName("value")
							if valueArg != nil {
								f.Dot("Default").Call(
									jen.Id(valueArg.Value.String()),
								)
							}
						default:
							continue
						}
					}
					g.Add(f)
				}
				g.Line()
			}),
		),
	)

	m.Line()

	m.Comment(fmt.Sprintf("Edges of the %s.", def.Name))
	m.Func().Parens(jen.Id(def.Name)).Id("Edges").Params().
		Index().Qual("github.com/facebook/ent", "Edge").Block(
		jen.Return(
			jen.Index().Qual("github.com/facebook/ent", "Edge").ValuesFunc(func(g *jen.Group) {
				for _, field := range forEdges {
					var e *jen.Statement = jen.Line()

					if field.Directives.ForName("from") != nil {
						e.Qual("github.com/facebook/ent/schema/edge", "From")
					} else if field.Directives.ForName("to") != nil {
						e.Qual("github.com/facebook/ent/schema/edge", "To")
					} else { // default
						e.Qual("github.com/facebook/ent/schema/edge", "To")
					}

					e.Call(
						jen.Lit(field.Name),
						jen.Id(field.Type.Name()).Dot("Type"),
					)

					if field.Type.Elem == nil { // unique
						e.Dot("Unique").Call()
					}

					if field.Type.NonNull {
						e.Dot("Required").Call()
					}

					refValue := ""

					if field.Directives.ForName("from") != nil {
						ref := field.Directives.ForName("from").Arguments.ForName("ref")
						if ref != nil {
							refValue = ref.Value.Raw
							e.Dot("Ref").Call(jen.Lit(refValue))
						}
					}

					g.Add(e)
				}
				g.Line()
			}),
		),
	)

	return m
}
