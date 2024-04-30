package schema

import (
	"context"
	"errors"
	"rr-backend/ent/entgen/hook"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type CommonAttribute struct {
	Id_ulid string
	Id_int  int
	mixin.Schema
}

func (CommonAttribute) Fields() []ent.Field {
	return []ent.Field{
		field.String("CreatedBy").StorageKey("CreatedBy").MaxLen(40).Optional().Nillable(),
		field.String("UpdatedBy").StorageKey("UpdatedBy").MaxLen(40).Optional().Nillable(),
		field.String("DeletedBy").StorageKey("DeletedBy").MaxLen(40).Optional().Nillable(),

		field.String("Ip").StorageKey("Ip").MaxLen(15).Optional().Nillable(),
		field.String("UserClient").StorageKey("UserClient").Optional().Nillable(),
		field.Time("CreatedAt").StorageKey("CreatedAt").Default(time.Now().UTC()),
		field.Time("UpdatedAt").StorageKey("UpdatedAt").Default(time.Now()).UpdateDefault(time.Now()),
		field.Time("DeletedAt").StorageKey("DeletedAt"),
	}
}

func (CommonAttribute) Hook() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {

			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				mc, ok := m.(interface {
					SetCreateBy(string)
				})

				if !ok {
					return nil, errors.New("unexpected mutation found")
				}

				userId := ctx.Value("UserId").(string)

				mc.SetCreateBy(userId)

				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),

		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(
					func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
						mc, ok := m.(interface {
							SetUpdatedBy(string)
						})

						if !ok {
							return nil, errors.New("unexpected mutation found")
						}

						userId := ctx.Value("UserId").(string)

						mc.SetUpdatedBy(userId)

						return next.Mutate(ctx, m)
					},
				)
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
