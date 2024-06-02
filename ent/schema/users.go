package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TblUSers struct {
	ent.Schema
}

func (TblUSers) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Tbl_Users"},
	}
}

func (TblUSers) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonAttribute{
			Id_ulid: "Id_ulid",
		},
	}
}

func (TblUSers) Fields() []ent.Field {
	return []ent.Field{
		field.String("UserName").StorageKey("UserName").MaxLen(40),
		field.String("Password").StorageKey("Password").MaxLen(40),
		field.String("Email").StorageKey("Email").MaxLen(40).Optional().Nillable(),
	}
}

func (TblUSers) Edges() []ent.Edge {
	return []ent.Edge{}
}
