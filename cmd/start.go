package cmd

import (
	"rr-backend/config"
	"rr-backend/config/env"
	"rr-backend/config/logx"
	"rr-backend/errorx"
	"rr-backend/routes"

	"github.com/labstack/echo/v4"
)

func Start() {
	env.LoadEnv()
	logx.InitLogger()
	config := config.NewAppConfig()

	RunMigration(config.GetSqlxClient().DB)

	e := echo.New()

	e.HTTPErrorHandler = errorx.CustomErrorHandler

	routes.InitRoutes(e, config)

	e.Logger.Fatal(e.Start(":8080"))
}
