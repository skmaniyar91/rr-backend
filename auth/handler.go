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
	storage IAuthStorage
}

func NewAuthHandler(config config.IAppConfig) IAuthHandler {
	return &sAuthHandler{
		storage: NewAuthStorage(config),
	}
}

func (s *sAuthHandler) Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "" || password == "" {
		return c.JSON(http.StatusBadRequest, "missing credentials")
	}

	token, err := s.storage.Login(username, password)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}

func (s *sAuthHandler) LoginTest(c echo.Context) error {
	id := c.Get("UserId")
	return c.JSON(http.StatusOK, fmt.Sprintf("yehh, your authentication works, %s", id.(string)))
}

func (s *sAuthHandler) Logout(c echo.Context) error {
	c.Set("UserId", "")

	return c.JSON(http.StatusOK, nil)
}
