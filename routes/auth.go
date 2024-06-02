package routes

import (
	"rr-backend/auth"
	"rr-backend/config"

	"github.com/labstack/echo/v4"
)

func initAuthRoutes(apiGrp *echo.Group, config config.IAppConfig) {
	authHandler := auth.NewAuthHandler(config)

	apiGrp.GET("/login", authHandler.Login)
	apiGrp.GET("/logout", authHandler.Logout)
	apiGrp.GET("/login-test", authHandler.LoginTest, auth.Auth)
}
