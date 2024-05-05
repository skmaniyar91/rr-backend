package restmdl

import "github.com/labstack/echo"

type RequestMetaData struct {
	UserId *string

	Ip        *string
	UserAgent *string
}

func GetRequestMetaData(c echo.Context) RequestMetaData {
	rmd := c.Get("rmd").(RequestMetaData)
	temp := c.RealIP()
	rmd.Ip = &temp
	temp = c.Request().UserAgent()
	rmd.UserAgent = &temp

	return rmd
}
