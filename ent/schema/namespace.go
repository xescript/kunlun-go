package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Namespace struct {
	ent.Schema
}

func (Namespace) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Namespace) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").
			MaxLen(512).
			Unique(),
		field.String("full_name").
			MaxLen(512).
			Default(""),
		field.Enum("type").
			Values("user", "group"),
	}
}

func (Namespace) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			StorageKey(edge.Column("namespace_id")).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
		edge.To("group", Group.Type).
			Unique().
			StorageKey(edge.Column("namespace_id")).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
