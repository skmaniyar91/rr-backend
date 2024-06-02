package routes

import (
	"rr-backend/config"

	"github.com/labstack/echo/v4"

	_ "rr-backend/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo, appcongif config.IAppConfig) {
	apiGrp := e.Group("/api/rr-backend")

	apiGrp.GET("/swagger/*", echoSwagger.WrapHandler)

	initAuthRoutes(apiGrp, appcongif)
	initiUsersRoutes(*apiGrp, appcongif)
}
