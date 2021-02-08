package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Autor  holds the schema definition for the Autor entity.
type Autor struct {
	ent.Schema
}

// Fields of the Autor.
func (Autor) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
		field.String("name"),
		field.String("signature").Optional().Unique(),
	}
}

// Edges of the Autor.
func (Autor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type),
	}
}
