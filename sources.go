package life

import "github.com/vektah/gqlparser/v2/ast"

const basics = `
directive @from(ref: String!) on FIELD_DEFINITION

directive @unique on FIELD_DEFINITION
directive @immutable on FIELD_DEFINITION
directive @nillable on FIELD_DEFINITION
directive @sensitive on FIELD_DEFINITION

directive @storageKey(key: String!) on FIELD_DEFINITION
directive @default(value: String!) on FIELD_DEFINITION
directive @updateDefault(value: String!) on FIELD_DEFINITION

directive @ent(name: String) on OBJECT

directive @out on OBJECT
`

func lifeDirectivesSource() *ast.Source {
	return &ast.Source{
		Input:   basics,
		Name:    "goent.graphql",
		BuiltIn: true,
	}
}

// DirectivesSource returns an graphql ast source of directives used in life for entgo binds.
func DirectivesSource() *ast.Source {
	return lifeDirectivesSource()
}
