package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique().Immutable(),
		field.String("title").NotEmpty(),
		field.String("body").NotEmpty(),
	}
}

// Edges of the Post.	
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("autor", Person.Type).Unique().Ref("posts"),
	
	}
}
