package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/iancoleman/strcase"

	"github.com/dave/jennifer/jen"
	"github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

var scalars = map[string]jen.Code{
	"ID":      jen.Id("ID"),
	"Int":     jen.Int(),
	"Float":   jen.Float64(),
	"Boolean": jen.Bool(),
	"String":  jen.String(),
}

var entScalars = map[string]*jen.Statement{
	"ID":      jen.Qual("github.com/facebook/ent/schema/field", "String"),
	"Int":     jen.Qual("github.com/facebook/ent/schema/field", "Int"),
	"Float":   jen.Qual("github.com/facebook/ent/schema/field", "Float64"),
	"Boolean": jen.Qual("github.com/facebook/ent/schema/field", "Boolean"),
	"String":  jen.Qual("github.com/facebook/ent/schema/field", "String"),
}

func openSchemaSource(filepath string) (*ast.Source, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	source := &ast.Source{
		Input:   string(data),
		Name:    filepath,
		BuiltIn: false,
	}

	return source, nil
}

func fieldTypeGen(t *ast.Type) jen.Code {
	fieldStatement := jen.Statement{}

	if !t.NonNull {
		fieldStatement = append(fieldStatement, jen.Op("*"))
	}

	if t.Elem != nil {
		fieldStatement = append(fieldStatement, jen.Index())
		if !t.Elem.NonNull {
			fieldStatement = append(fieldStatement, jen.Op("*"))
		}
	}

	fieldStatement = append(fieldStatement, jen.Id(t.Name()))

	return &fieldStatement
}

func fieldScalarGen(field *ast.FieldDefinition, fieldType jen.Code) jen.Code {
	// spew.Dump(field)

	// if field.Type.Elem != nil { // is array
	// 	// switch field.Type.Name() {
	// 	// case "ID":
	// 	// }
	// }

	name := strcase.ToCamel(field.Name)

	fieldStatement := jen.Statement{}

	fieldStatement = append(fieldStatement, jen.Id(name))

	if !field.Type.NonNull {
		fieldStatement = append(fieldStatement, jen.Op("*"))
	}

	if field.Type.Elem != nil {
		fieldStatement = append(fieldStatement, jen.Index())
		if !field.Type.Elem.NonNull {
			fieldStatement = append(fieldStatement, jen.Op("*"))
		}
	}

	// if isKnowScalar {
	fieldStatement = append(fieldStatement, fieldType)
	// } else { // probably a non scalar type
	// 	fieldStatement = append(fieldStatement, jen.Id(field.Type.Name()))
	// }

	return &fieldStatement
}

func generateType(def *ast.Definition) *jen.File {
	m := jen.NewFile("schema")

	m.ImportName("github.com/facebook/ent", "ent")
	m.ImportName("github.com/facebook/ent/schema/field", "field")
	m.ImportName("github.com/facebook/ent/schema/edge", "edge")

	m.Type().Id(def.Name).Struct(
		jen.Qual("github.com/facebook/ent", "Schema"),
	)

	log.Println("---- " + def.Name)

	forFields := []*ast.FieldDefinition{}
	forEdges := []*ast.FieldDefinition{}

	for _, field := range def.Fields {
		_, isKnownScalar := entScalars[field.Type.Name()]
		if isKnownScalar {
			forFields = append(forFields, field)
		} else {
			forEdges = append(forEdges, field)
		}
	}

	m.Func().Parens(jen.Id(def.Name)).Id("Fields").Params().
		Index().Qual("github.com/facebook/ent", "Field").Block(
		jen.Return(
			jen.Index().Qual("github.com/facebook/ent", "Field").ValuesFunc(func(g *jen.Group) {
				for _, field := range forFields {
					fieldScalar := entScalars[field.Type.Name()]

					log.Println("- field ", field.Name, field.Type.Name())

					f := fieldScalar.Call(jen.Lit(field.Name))

					if field.Type.NonNull {
						f.Dot("NotEmpty").Call()
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
					// g.Comment("str string")
					g.Add(f)
				}
			}),
		),
	)

	m.Line()

	m.Func().Parens(jen.Id(def.Name)).Id("Edges").Params().
		Index().Qual("github.com/facebook/ent", "Edge").Block(
		jen.Return(
			jen.Index().Qual("github.com/facebook/ent", "Edge").ValuesFunc(func(g *jen.Group) {
				for _, field := range forEdges {
					log.Println("- edge ", field.Name, field.Type.Name())
					var e *jen.Statement

					if field.Directives.ForName("from") != nil {
						e = g.Qual("github.com/facebook/ent/schema/edge", "From")
					} else if field.Directives.ForName("to") != nil {
						e = g.Qual("github.com/facebook/ent/schema/edge", "To")
					} else { // default
						e = g.Qual("github.com/facebook/ent/schema/edge", "To")
					}

					e.Call(
						jen.Lit(field.Name),
						jen.Id(field.Type.Name()).Dot("Type"),
					)

					refValue := ""

					if field.Directives.ForName("from") != nil {
						ref := field.Directives.ForName("from").Arguments.ForName("ref")
						if ref != nil {
							refValue = ref.Value.String()
							e.Dot("Ref").Call(jen.Lit(refValue))
						}
					}

					if field.Type.Elem == nil { // unique
						e.Dot("Unique").Call()
					}

					if field.Type.NonNull {
						e.Dot("Required").Call()
					}

					g.Add(e)
				}

				g.Line()
			}),
		),
	)

	return m
}

func generateTypes(folder string, definitions []*ast.Definition) {
	for _, def := range definitions {
		entitityFile := generateType(def)

		filepath := path.Join(folder, strcase.ToLowerCamel(def.Name)+".go")
		f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0660)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if err = entitityFile.Render(f); err != nil {
			panic(err)
		}
	}
}

func generate(filepath string, folderOut string) {
	source, err := openSchemaSource(filepath)
	if err != nil {
		panic(err)
	}

	sch, gqlErr := gqlparser.LoadSchema(source)
	if gqlErr != nil {
		log.Println(gqlErr.Message)
		panic(err)
	}

	defs := []*ast.Definition{}
	queries := []*ast.Definition{}
	mutations := []*ast.Definition{}

	for _, t := range sch.Types {
		if t.Name == "Query" { // make something with mutations and queries
			queries = append(queries, t)
		} else if t.Name == "Mutation" {
			mutations = append(mutations, t)
		} else if t.Name == "Subscription" {
			continue
		} else if t.BuiltIn {
			continue
		} else if t.Kind == ast.InputObject {
			continue
		} else if t.Kind == ast.Enum {
			continue
		} else if t.Kind == ast.Interface {
			continue
		} else if t.Kind == ast.Scalar {
			continue
		} else if t.Kind == ast.Union {
			continue
		} else {
			defs = append(defs, t)
		}
	}

	// for _, def := range defs {
	// 	log.Println(def.Name)
	// }

	generateTypes(folderOut, defs)
}

func main() {
	// strcase.ConfigureAcronym
	generate("example/schema.medium.graphql", "schema")
}
