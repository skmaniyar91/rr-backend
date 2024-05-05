package users

import (
	"net/http"
	"rr-backend/config"
	"rr-backend/lib"

	"github.com/labstack/echo"
)

type IUserHandler interface {
	CreateUser(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

type sUserHandler struct {
	userService IUSerService
}

func NewUserHandler(config config.IAppConfig) IUserHandler {
	return &sUserHandler{
		userService: NewUserService(config),
	}
}

// CreateUser implements IUserHandler.
func (s *sUserHandler) CreateUser(c echo.Context) error {
	var rq RQUser

	err := c.Bind(&rq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "binding error")
	}

	rq.rmd = lib.GetRequestMetaData(c)

	rs, err := s.userService.CreateUser(rq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, rs)
}

// DeleteUser implements IUserHandler.
func (s *sUserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	rmd := lib.GetRequestMetaData(c)

	err := s.userService.DeleteUser(id, rmd)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, nil)
}

// GetUser implements IUserHandler.
func (s *sUserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")

	rs, err := s.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, rs)
}

// UpdateUser implements IUserHandler.
func (s *sUserHandler) UpdateUser(c echo.Context) error {
	var rq RQUser

	rq.Id = c.Param("id")
	c.Bind(&rq)

	rq.rmd = lib.GetRequestMetaData(c)

	rs, err := s.userService.UpdateUser(rq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, rs)
}
