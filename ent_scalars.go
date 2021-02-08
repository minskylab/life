package life

import "github.com/dave/jennifer/jen"

var entScalars = map[string]*jen.Statement{
	"ID":       jen.Qual("entgo.io/ent/schema/field", "String"),
	"Int":      jen.Qual("entgo.io/ent/schema/field", "Int64"),
	"Float":    jen.Qual("entgo.io/ent/schema/field", "Float"),
	"Boolean":  jen.Qual("entgo.io/ent/schema/field", "Bool"),
	"String":   jen.Qual("entgo.io/ent/schema/field", "String"),
	"Map":      jen.Qual("entgo.io/ent/schema/field", "JSON"),
	"Time":     jen.Qual("entgo.io/ent/schema/field", "Time"),
	"DateTime": jen.Qual("entgo.io/ent/schema/field", "Time"),
	"Enum":     jen.Qual("entgo.io/ent/schema/field", "Enum"),
}

var entScalarArrays = map[string]*jen.Statement{
	"Int":    jen.Qual("entgo.io/ent/schema/field", "Ints"),
	"Float":  jen.Qual("entgo.io/ent/schema/field", "Floats"),
	"String": jen.Qual("entgo.io/ent/schema/field", "Strings"),
}
