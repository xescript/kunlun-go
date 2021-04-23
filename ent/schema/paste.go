package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Paste struct {
	ent.Schema
}

func (Paste) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Paste) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			MaxLen(512),
		field.Text("content").
			NotEmpty().
			SchemaType(map[string]string{
				dialect.SQLite: "TEXT",
				dialect.MySQL:  "LONGTEXT",
			}),
	}
}

func (Paste) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).
			StorageKey(edge.Column("owner_id")).
			Unique().
			Required().
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
	}
}
