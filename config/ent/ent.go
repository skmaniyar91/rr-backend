package ent

import (
	"rr-backend/ent/entgen"
	_ "rr-backend/ent/entgen/runtime"
	constant "rr-backend/lib/const"
)

func GetENTClient() *entgen.Client {
	client, err := entgen.Open("mysql", constant.DatabaseURL)
	if err != nil {
		panic(err)
	}

	return client
}
