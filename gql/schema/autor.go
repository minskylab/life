package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

type Autor struct {
	ent.Schema
}

func (Autor) Fields() []ent.Field {
	return []ent.Field{field.String("id").NotEmpty(), field.String("name").NotEmpty()}
}

func (Autor) Edges() []ent.Edge {
	return []ent.Edge{edge.To("todos", Todo.Type).Required(), edge.To("todos", Todo.Type).Required()}
}
