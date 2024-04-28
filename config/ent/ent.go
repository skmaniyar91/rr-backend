package ent

import (
	"log"
	"rr-backend/ent/entgen"
	constant "rr-backend/lib/const"
)

func GetENTClient() *entgen.Client {
	client, err := entgen.Open("mysql", constant.DatabaseURL+"parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	defer client.Close()

	return client
}
