package auth

import (
	"fmt"
	"net/http"
	"rr-backend/config"

	"github.com/labstack/echo/v4"
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

// Login
//
//	@Summary		Login
//	@Description	Login
//	@Tags			Auth API
//	@Accept			mpfd
//	@Produce		json
//	@Param			username	formData	string	true	"UserName"
//	@Param			password	formData	string	true	"Password"
//	@Success		200			{object}	RSToken
//	@Failure		500			{string}	string
//	@Router			/api/rr-backend/login [post]
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

// LoginTest
//
//	@Summary		Login Test
//	@Description	Login Test
//	@Tags			Auth API
//	@Accept			mpfd
//	@Produce		json
//	@Security		JWTAuthentication
//	@Success		200	{string}	string
//	@Failure		500	{string}	string
//	@Router			/api/rr-backend/login-test [get]
func (s *sAuthHandler) LoginTest(c echo.Context) error {
	id := c.Get("UserId")
	return c.JSON(http.StatusOK, fmt.Sprintf("yehh, your authentication works, %s", id.(string)))
}

// Login
//
//	@Summary		Logout
//	@Description	Logout
//	@Tags			Auth API
//	@Success		200
//	@Failure		500	{string}	string
//	@Router			/api/rr-backend/logout [get]
func (s *sAuthHandler) Logout(c echo.Context) error {
	c.Set("UserId", "")

	return c.JSON(http.StatusOK, nil)
}
