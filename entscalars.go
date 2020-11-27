package life

import "github.com/dave/jennifer/jen"

var entScalars = map[string]*jen.Statement{
	"ID":       jen.Qual("github.com/facebook/ent/schema/field", "String"),
	"Int":      jen.Qual("github.com/facebook/ent/schema/field", "Int"),
	"Float":    jen.Qual("github.com/facebook/ent/schema/field", "Float64"),
	"Boolean":  jen.Qual("github.com/facebook/ent/schema/field", "Bool"),
	"String":   jen.Qual("github.com/facebook/ent/schema/field", "String"),
	"Map":      jen.Qual("github.com/facebook/ent/schema/field", "JSON"),
	"Time":     jen.Qual("github.com/facebook/ent/schema/field", "Time"),
	"DateTime": jen.Qual("github.com/facebook/ent/schema/field", "Time"),
}
