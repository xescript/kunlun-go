package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

type Group struct {
	ent.Schema
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Group) Fields() []ent.Field {
	return []ent.Field{
	}
}

func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("namespace", Namespace.Type).
			Ref("group").
			Unique().
			Required(),
	}
}
