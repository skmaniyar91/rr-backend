package auth

import (
	"net/http"

	"github.com/labstack/echo"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")

		if tokenStr == "" {
			return c.JSON(http.StatusBadRequest, "token not found")
		}
		tokenStr = tokenStr[len("Bearer "):]
		if tokenStr == "" {
			return c.JSON(http.StatusBadRequest, "token not found")
		}

		claims, err := ValidateToken(tokenStr)
		if err != nil {
			return err
		}

		c.Set("UserId", claims.UserId)

		return next(c)
	}
}
