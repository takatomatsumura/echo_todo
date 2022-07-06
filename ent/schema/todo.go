package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("content"),
		field.Bool("todoComplete").Default(false),
		field.Time("deadline").Optional(),
		field.Time("created_at").Default(time.Now).Optional(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Optional(),
		field.String("image_path").Optional(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("todos"),
	}
}
