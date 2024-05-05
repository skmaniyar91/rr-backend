package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type TblSuperAdmin struct {
	ent.Schema
}

func (TblSuperAdmin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Tbl_SuperAdmin"},
	}
}

func (TblSuperAdmin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonAttribute{
			Id_ulid: "Id_ulid",
		},
	}
}

func (TblSuperAdmin) Fields() []ent.Field {
	return []ent.Field{
		field.String("UserName").StorageKey("UserName").MaxLen(40),
		field.String("PassWord").StorageKey("PassWord").MaxLen(40),
	}
}

func (TblSuperAdmin) Edges() []ent.Edge {
	return nil
}
