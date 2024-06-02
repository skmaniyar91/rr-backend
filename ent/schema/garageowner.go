package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type TblGarageOwner struct {
	ent.Schema
}

func (TblGarageOwner) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Tbl_GarageOwner",
		},
	}
}

func (TblGarageOwner) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonAttribute{
			Id_ulid: "Id_ulid",
		},
	}
}

func (TblGarageOwner) Fields() []ent.Field {
	return []ent.Field{
		field.String("UserId_ulid").StorageKey("UserId_ulid").Optional().Nillable().MaxLen(40),
		field.Int("Initial").StorageKey("Initial").Optional().Nillable(),

		field.String("FirstName").StorageKey("FirstName").MaxLen(100),
		field.String("MiddleName").StorageKey("MiddleName").MaxLen(100).Optional().Nillable(),
		field.String("LastName").StorageKey("LastName").MaxLen(100),

		field.String("ContactNumber").StorageKey("ContactNumber").MaxLen(50),
		field.String("Email").StorageKey("Email").MaxLen(100).Nillable().Optional(),

		field.Int("Age").StorageKey("Age").Optional().Nillable(),

		field.String("PhotoId_ulid").StorageKey("Photo").MaxLen(40).Optional().Nillable(),

		field.String("AddressId_ulid").StorageKey("Address").MaxLen(40),
	}
}

func (TblGarageOwner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", TblUSers.Type).Ref("Owner").Field("UserId_ulid").Unique(),
		edge.From("NameInitial", TblEnum.Type).Ref("InitialEnum").Field("Initial").Unique(),

		edge.From("OwnerPhoto", TblDocument.Type).Ref("Photo").Field("PhotoId_ulid").Unique(),
		edge.From("Address", TblAddress.Type).Required().Ref("OwnerAddress").Field("AddressId_ulid").Unique(),
	}
}
