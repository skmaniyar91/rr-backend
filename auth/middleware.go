package auth

import (
	"rr-backend/errorx"

	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenStr := c.Request().Header.Get("Authorization")

		if tokenStr == "" {
			return errorx.NewUnProccessableEntity("auth", "Token not found")
		}
		tokenStr = tokenStr[len("Bearer "):]
		if tokenStr == "" {
			return errorx.NewUnProccessableEntity("auth", "Token not found")
		}

		claims, err := ValidateToken(tokenStr)
		if err != nil {
			return errorx.NewUnauthorizedError("auth", "Invalid token")
		}

		c.Set("UserId", claims.UserId)

		return next(c)
	}
}
