package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_admin").
			Default(false).
			Comment("Super admin user"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("namespace", Namespace.Type).
			Ref("user").
			Unique().
			Required(),
	}
}
