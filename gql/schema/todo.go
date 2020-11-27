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
	return []ent.Field{field.String("id").NotEmpty(), field.String("title").NotEmpty(), field.Bool("done")}
}

func (Todo) Edges() []ent.Edge {
	return []ent.Edge{edge.To("owner", Autor.Type).Unique()}
}
