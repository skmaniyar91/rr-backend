package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TblDocument struct {
	ent.Schema
}

func (TblDocument) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonAttribute{
			Id_ulid: "Id_ulid",
		},
	}
}

func (TblDocument) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Tbl_Document",
		},
	}
}

func (TblDocument) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name").StorageKey("Name").MaxLen(50),
		field.String("RelativePath").StorageKey("RelativePath").MaxLen(500),
		field.String("URL").StorageKey("URL").MaxLen(1000),
		field.Float("SizeInBytes").StorageKey("SizeInBytes"),
	}
}

func (TblDocument) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Photo", TblGarageOwner.Type).Unique(),
	}
}
