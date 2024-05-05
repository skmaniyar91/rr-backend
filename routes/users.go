package routes

import (
	"rr-backend/config"
	"rr-backend/internal/users"
	constant "rr-backend/lib/const"

	"github.com/labstack/echo"
)

func initiUsersRoutes(apiGrp echo.Group, config config.IAppConfig) {
	usersHandler := users.NewUserHandler(config)

	userGrp := apiGrp.Group(constant.RouteUser)

	userGrp.POST("", usersHandler.CreateUser)

	userGrp.GET("/:id", usersHandler.GetUser)
	userGrp.PUT("/:id", usersHandler.UpdateUser)
	userGrp.DELETE("/:id", usersHandler.DeleteUser)
}
