package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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

func fieldGen(field *ast.FieldDefinition) jen.Code {
	// spew.Dump(field)

	// if field.Type.Elem != nil { // is array
	// 	// switch field.Type.Name() {
	// 	// case "ID":
	// 	// }
	// }

	name := strings.ToUpper(string(field.Name[0])) + string(field.Name[1:])

	fieldStatement := jen.Statement{}

	fieldStatement = append(fieldStatement, jen.Id(name))

	fieldType, isKnowScalar := scalars[field.Type.Name()]

	if !field.Type.NonNull {
		fieldStatement = append(fieldStatement, jen.Op("*"))
	}

	if field.Type.Elem != nil {
		fieldStatement = append(fieldStatement, jen.Index())
		if !field.Type.Elem.NonNull {
			fieldStatement = append(fieldStatement, jen.Op("*"))
		}
	}

	if isKnowScalar {
		fieldStatement = append(fieldStatement, fieldType)
	} else { // probably a non scalar type
		fieldStatement = append(fieldStatement, jen.Id(field.Type.Name()))
	}

	return &fieldStatement
}

func generateTypes(definitions []*ast.Definition) {
	m := jen.NewFile("models")

	m.Type().Id("ID").String().Line()

	for _, def := range definitions {
		fields := []jen.Code{}
		for _, field := range def.Fields {
			fields = append(fields, fieldGen(field))
		}

		m.Type().Id(def.Name).Struct(fields...).Line()
	}

	fmt.Printf("%#v", m)
}

func parse(filepath string) {
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

		if !t.BuiltIn && !t.OneOf("Query", "Mutation") {
			defs = append(defs, t)
		}
	}

	generateTypes(defs)

}

func main() {
	parse("example/schema.graphql")
}
