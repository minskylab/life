package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Book  holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
		field.Time("createdAt"),
		field.Time("updatedAt"),
		field.Enum("kind").Values("NOVEL", "ESSAY", "JOURNAL"),
		field.String("doi").Default("2"),
		field.String("title"),
		field.String("description"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Autor.Type).Unique().Required().Ref("books"),
	}
}
