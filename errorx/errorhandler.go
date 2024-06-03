package errorx

import (
	"net/http"
	"rr-backend/lib/restmdl"

	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, c echo.Context) {
	baseError := HandleError(err)

	baseRS := restmdl.BaseRS{
		Error: baseError,
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(baseError.StatusCode)
		} else {
			err = c.JSON(baseError.StatusCode, baseRS)
		}
		if err != nil {
			c.Echo().Logger.Error(baseError)
		}
	}
}
