package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TblAddress struct {
	ent.Schema
}

func (TblAddress) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonAttribute{
			Id_ulid: "Id_ulid",
		},
	}
}

func (TblAddress) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Tbl_Address",
		},
	}
}

func (TblAddress) Fields() []ent.Field {
	return []ent.Field{
		field.String("Line1").StorageKey("Line1").MaxLen(50),
		field.String("Line2").StorageKey("Line2").MaxLen(50).Nillable().Optional(),
		field.String("Line3").StorageKey("Line3").MaxLen(50).Nillable().Optional(),

		field.String("City").StorageKey("City").MaxLen(50),
		field.String("District").StorageKey("District").MaxLen(50).Nillable().Optional(),
		field.String("SubDistrict").StorageKey("SubDistrict").MaxLen(50).Nillable().Optional(),

		field.String("State").StorageKey("State").MaxLen(50),
		field.String("Country").StorageKey("Country").MaxLen(50),

		field.String("PostalCode").StorageKey("PostalCode").MaxLen(50).Nillable().Optional(),
	}
}

func (TblAddress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("OwnerAddress", TblGarageOwner.Type).Unique(),
	}
}
