package cmd

import (
	"rr-backend/config"
	"rr-backend/config/env"
)

func Start() {
	env.LoadEnv()

	config := config.NewAppConfig()

	RunMigration(config.GetSqlxClient().DB)
}
