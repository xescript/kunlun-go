package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TransactionMixin struct {
	mixin.Schema
}

func (TransactionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").
			NotEmpty().
			MaxLen(36),
		field.Text("old_value").
			Nillable().
			Optional().
			SchemaType(map[string]string{
				dialect.SQLite: "TEXT",
				dialect.MySQL:  "LONGTEXT",
			}),
		field.Text("new_value").
			Nillable().
			Optional().
			SchemaType(map[string]string{
				dialect.SQLite: "TEXT",
				dialect.MySQL:  "LONGTEXT",
			}),
	}
}

func (TransactionMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).
			StorageKey(edge.Column("author_id")).
			Unique().
			Required().
			Annotations(entsql.Annotation{
				OnDelete: entsql.SetNull,
			}),
	}
}
