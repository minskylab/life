package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
		field.String("name").NotEmpty(),
		field.String("email").NotEmpty(),
	}
}

// Edges of the Person.	
func (Person) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}
