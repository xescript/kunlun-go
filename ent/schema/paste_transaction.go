package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
)

type PasteTransaction struct {
	ent.Schema
}

func (PasteTransaction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TransactionMixin{},
		TimeMixin{},
	}
}
func (PasteTransaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("paste", Paste.Type).
			StorageKey(edge.Column("paste_id")).
			Unique().
			Required().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
