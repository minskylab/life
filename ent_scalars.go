package life

import "github.com/dave/jennifer/jen"

var entScalars = map[string]*jen.Statement{
	"ID":       jen.Qual(entBasePath+"/schema/field", "String"),
	"Int":      jen.Qual(entBasePath+"/schema/field", "Int64"),
	"Float":    jen.Qual(entBasePath+"/schema/field", "Float"),
	"Boolean":  jen.Qual(entBasePath+"/schema/field", "Bool"),
	"String":   jen.Qual(entBasePath+"/schema/field", "String"),
	"Map":      jen.Qual(entBasePath+"/schema/field", "JSON"),
	"Time":     jen.Qual(entBasePath+"/schema/field", "Time"),
	"DateTime": jen.Qual(entBasePath+"/schema/field", "Time"),
	"Enum":     jen.Qual(entBasePath+"/schema/field", "Enum"),
}

var entScalarArrays = map[string]*jen.Statement{
	"Int":    jen.Qual(entBasePath+"/schema/field", "Ints"),
	"Float":  jen.Qual(entBasePath+"/schema/field", "Floats"),
	"String": jen.Qual(entBasePath+"/schema/field", "Strings"),
}
