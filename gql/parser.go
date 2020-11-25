package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

func generateTypes(definitions []*ast.Definition) *jen.File {
	m := jen.NewFile("models")

	m.Type().Id("ID").String().Line()
	m.Type().Id("DateTime").Qual("time", "Time").Line()

	for _, def := range definitions {
		if def.Kind == ast.Enum {
			m.Type().Id(def.Name).String().Line()
			for _, val := range def.EnumValues {
				// log.Println(val.Name)
				m.Const().Id(strcase.ToScreamingSnake(val.Name)).Id(def.Name).Op("=").Lit(val.Name)
			}
		} else if def.Kind == ast.Scalar {
			continue
		} else {
			fields := []jen.Code{}
			queryFields := []*ast.FieldDefinition{}

			for _, field := range def.Fields {
				fieldType, isKnowScalar := scalars[field.Type.Name()]

				if isKnowScalar {
					fields = append(fields, fieldScalarGen(field, fieldType))
				} else {
					queryFields = append(queryFields, field)
				}
			}

			if def.IsInputType() {
				for _, field := range queryFields {
					fieldType := fieldTypeGen(field.Type)
					fields = append(fields, &jen.Statement{jen.Id(field.Name), fieldType})
				}
			}

			// root struct type
			m.Comment(fmt.Sprintf("%s is kind %s", def.Name, string(def.Kind)))
			m.Type().Id(def.Name).Struct(fields...).Line()

			// dynamics

			if !def.IsInputType() {
				for _, field := range queryFields {
					args := []jen.Code{}
					for _, arg := range field.Arguments {
						fieldType, isKnowScalar := scalars[arg.Type.Name()]

						if isKnowScalar {
							sts := jen.Statement{jen.Id(arg.Name)}
							if !arg.Type.NonNull {
								sts = append(sts, jen.Op("*"))
							}
							sts = append(sts, fieldType)
							args = append(args, &sts)
						} else {
							t := fieldTypeGen(arg.Type)
							args = append(args, &jen.Statement{jen.Id(arg.Name), t})
						}

					}
					name := strcase.ToCamel(field.Name)
					m.Func().Params(
						jen.Id(strcase.ToLowerCamel(def.Name)).Op("*").Id(def.Name),
					).Id(name).Params(
						args...,
					).Parens(jen.List(
						fieldTypeGen(field.Type),
						jen.Error(),
					)).Block(
						// jen.Comment("unimplemented"),
						jen.Panic(jen.Lit("unimplemented builder")),
						// jen.Return(jen.Nil(), jen.Nil()),
					).Line()
				}
			}
		}

	}

	return m
}

func generate(filepath string) {
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
	for _, t := range sch.Types {
		if t.OneOf("Query", "Mutation") {
			// make something with mutations and queries
		} else if !t.BuiltIn {
			defs = append(defs, t)
		}
	}

	models := generateTypes(defs)
	f, err := os.OpenFile("models/generated.go", os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if err = models.Render(f); err != nil {
		panic(err)
	}
}

func main() {
	// strcase.ConfigureAcronym
	generate("example/schema.graphql")
}
