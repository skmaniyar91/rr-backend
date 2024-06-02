package ent

import (
	"context"
	"rr-backend/ent/entgen"
	"rr-backend/ent/entgen/intercept"
	_ "rr-backend/ent/entgen/runtime"
	constant "rr-backend/lib/const"

	"entgo.io/ent/dialect/sql"
)

func GetENTClient() *entgen.Client {
	client, err := entgen.Open("mysql", constant.DatabaseURL+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	client.Intercept(
		intercept.Func(func(ctx context.Context, q intercept.Query) error {
			q.WhereP(func(s *sql.Selector) {
				sql.FieldIsNull("DeletedAt")
			})
			return nil
		}),
	)

	return client
}
