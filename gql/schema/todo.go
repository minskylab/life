package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

type Todo struct {
	ent.Schema
}

func (Todo) Fields() []ent.Field {
	return []ent.Field{field.String("id").NotEmpty()("id").NotEmpty(), field.String("name").NotEmpty()("title").NotEmpty(), field.Boolean("done").NotEmpty()}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{edge.To("owner", Autor.Type).Unique(), edge.To("owner", Autor.Type).Unique()}
}
