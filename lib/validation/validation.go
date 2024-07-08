package validation

import (
	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindAndValiate(c echo.Context, model interface{}) error {
	c.Bind(model)

	validate := validator.New()
	err := validate.Struct(model)
	return err
}
