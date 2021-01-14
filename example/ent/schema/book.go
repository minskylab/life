package schema

import (
	"time"

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
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").UpdateDefault(time.Now),
		field.Float("doi").Default(2.32),
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
