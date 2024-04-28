package routes

import (
	"rr-backend/config"

	"github.com/labstack/echo"
)

func InitRoutes(e *echo.Echo, appcongif config.IAppConfig) {
	_ = e.Group("/api/rr-backend")
}
