package routes

import (
	"rr-backend/config"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo, appcongif config.IAppConfig) {
	apiGrp := e.Group("/api/rr-backend")

	initAuthRoutes(apiGrp, appcongif)
	initiUsersRoutes(*apiGrp, appcongif)
}
