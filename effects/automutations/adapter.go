package automutations

import (
	pluralize "github.com/gertd/go-pluralize"
	"github.com/vektah/gqlparser/v2/ast"
)

type entityField struct {
	Type       string
	Required   bool
	IsMultiple bool
}

type entityStructure struct {
	Name       string
	PluralName string

	ScalarFields   map[string]entityField
	RelationFields map[string]entityField
}

var scalars = []string{"ID", "Int", "Float", "Boolean", "String", "Map", "Time", "DateTime", "Enum"}

func (effect *EmergentEffect) isScalar(typeName string) bool {
	for _, s := range scalars {
		if s == typeName {
			return true
		}
	}
	return false
}

func (effect *EmergentEffect) generateStructure(types map[string]*ast.Definition, entity *ast.Definition) *entityStructure {
	pluralize := pluralize.NewClient()

	structure := &entityStructure{
		Name:           entity.Name,
		PluralName:     pluralize.Plural(entity.Name),
		ScalarFields:   map[string]entityField{},
		RelationFields: map[string]entityField{},
	}

	for _, field := range entity.Fields {
		name := field.Type.Name()

		if field.Name == "id" {
			continue
		}

		isMultiple := false
		if field.Type.Elem != nil {
			isMultiple = true
		}

		entField := entityField{
			Type:       name,
			Required:   field.Type.NonNull,
			IsMultiple: isMultiple,
		}

		if effect.isScalar(name) || types[field.Type.Name()].Kind == ast.Enum {
			structure.ScalarFields[field.Name] = entField
		} else {
			structure.RelationFields[field.Name] = entField
		}
	}

	return structure
}
