package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TblEnum struct {
	ent.Schema
}

func (TblEnum) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonAttribute{
			Id_int: "Id_int",
		},
	}
}

func (TblEnum) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Tbl_Enum",
		},
	}
}

func (TblEnum) Fields() []ent.Field {
	return []ent.Field{
		field.String("Code").StorageKey("Code").MaxLen(100),
		field.String("CodeType").StorageKey("CodeType").MaxLen(100),

		field.String("DisplayText").StorageKey("DisplayText").MaxLen(100),
	}
}

func (TblEnum) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("InitialEnum", TblGarageOwner.Type).Unique(),
	}
}
