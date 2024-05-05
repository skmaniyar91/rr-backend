package auth

import (
	"fmt"
	"net/http"
	"rr-backend/config"

	"github.com/labstack/echo"
)

type IAuthHandler interface {
	Login(c echo.Context) error
	LoginTest(c echo.Context) error
	Logout(c echo.Context) error
}

type sAuthHandler struct {
}

func NewAuthHandler(config config.IAppConfig) IAuthHandler {
	return &sAuthHandler{}
}

func (s *sAuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, "missing credentials")
	}

	if username == "samir" && password == "123" {
		accessToken, err := CreateToken("SM001")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		var token RSToken
		token.AccessToken = accessToken

		return c.JSON(http.StatusOK, token)
	}

	return c.JSON(http.StatusUnauthorized, "Unsuthorized")
}

func (s *sAuthHandler) LoginTest(c echo.Context) error {
	id := c.Get("UserId")
	return c.JSON(http.StatusOK, fmt.Sprintf("yehh, your authentication works, %s", id.(string)))
}

func (s *sAuthHandler) Logout(c echo.Context) error {
	c.Set("UserId", "")

	return c.JSON(http.StatusOK, nil)
}
