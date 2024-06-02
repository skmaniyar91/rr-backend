package schema

import (
	"context"
	"fmt"

	"rr-backend/ent/entgen"
	"rr-backend/ent/entgen/hook"
	"rr-backend/ent/entgen/intercept"
	"rr-backend/lib/restmdl"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
)

type CommonAttribute struct {
	Id_ulid string
	Id_int  string
	mixin.Schema
}

var generateUlid = func() string {
	return ulid.Make().String()
}

func (c CommonAttribute) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("CreatedBy").StorageKey("CreatedBy").MaxLen(40).Optional().Nillable(),
		field.String("UpdatedBy").StorageKey("UpdatedBy").MaxLen(40).Optional().Nillable(),
		field.String("DeletedBy").StorageKey("DeletedBy").MaxLen(40).Optional().Nillable(),

		field.String("IP").StorageKey("IP").Optional().Nillable(),
		field.String("UserAgent").StorageKey("UserAgent").Optional().Nillable(),
		field.Time("CreatedAt").StorageKey("CreatedAt").Default(time.Now),
		field.Time("UpdatedAt").StorageKey("UpdatedAt").Default(time.Now).UpdateDefault(time.Now),
		field.Time("DeletedAt").StorageKey("DeletedAt").Optional().Nillable(),
	}

	if c.Id_int != "" && c.Id_ulid != "" {
		fields = append(fields, field.String("id").StorageKey(c.Id_ulid).MaxLen(40).Immutable().DefaultFunc(generateUlid))
	} else if c.Id_ulid != "" {
		fields = append(fields, field.String("id").StorageKey(c.Id_ulid).MaxLen(40).Immutable().DefaultFunc(generateUlid))
	} else {
		fields = append(fields, field.Int("id").StorageKey(c.Id_ulid).Immutable())
	}

	return fields
}

func (c CommonAttribute) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value("SkipSoftDelete").(bool); skip {
				return nil
			}
			c.P(q)
			return nil
		}),
	}
}

func (CommonAttribute) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull("DeletedAt"),
	)
}

func (c CommonAttribute) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {

			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				mc, ok := m.(interface {
					SetCreatedBy(string)
				})

				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T from creating", m)
				}

				rmd := ctx.Value("rmd").(restmdl.RequestMetaData)
				mc.SetCreatedBy(*rmd.UserId)

				return next.Mutate(ctx, m)
			})
		}, ent.OpCreate),

		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(
					func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
						mx, ok := m.(interface {
							SetUpdatedBy(string)
							WhereP(...func(*sql.Selector))
						})

						if !ok {
							return nil, fmt.Errorf("unexpected mutation type %T from updating", m)
						}

						rmd := ctx.Value("rmd").(restmdl.RequestMetaData)
						c.P(mx)
						mx.SetUpdatedBy(*rmd.UserId)
						return next.Mutate(ctx, m)
					},
				)
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),

		hook.On(func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				// skip soft delete means delete entry permanantly
				if skip, _ := ctx.Value("SkipSoftDelete").(bool); skip {
					return next.Mutate(ctx, m)
				}

				mx, ok := m.(interface {
					SetOp(ent.Op)
					SetDeletedAt(time.Time)
					Client() *entgen.Client
					SetDeletedBy(string)
					WhereP(...func(*sql.Selector))
				})
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T deleting", m)
				}

				rmd := ctx.Value("rmd").(restmdl.RequestMetaData)

				c.P(mx)
				mx.SetOp(ent.OpUpdate)
				mx.SetDeletedAt(time.Now())
				mx.SetDeletedBy(*rmd.UserId)

				return mx.Client().Mutate(ctx, m)
			})
		}, ent.OpDelete|ent.OpDeleteOne),

		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				mx, ok := m.(interface {
					SetIP(string)
					SetUserAgent(string)
					WhereP(...func(*sql.Selector))
				})
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T common", m)
				}

				rmd := ctx.Value("rmd").(restmdl.RequestMetaData)
				c.P(mx)
				mx.SetIP(*rmd.Ip)
				mx.SetUserAgent(*rmd.UserAgent)

				return next.Mutate(ctx, m)
			})
		},
	}
}
