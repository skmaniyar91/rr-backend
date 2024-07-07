package lib

import (
	"context"
	"rr-backend/ent/entgen"
	"rr-backend/errorx"
	"rr-backend/lib/restmdl"

	"github.com/labstack/echo/v4"
)

func GetRequestMetaData(c echo.Context) restmdl.RequestMetaData {
	var rmd restmdl.RequestMetaData

	temp := c.RealIP()
	rmd.Ip = &temp

	temp = c.Request().UserAgent()
	rmd.UserAgent = &temp

	temp, _ = c.Get("UserId").(string)
	rmd.UserId = &temp

	return rmd
}

func GetContextWithRequestMetadata(rmd restmdl.RequestMetaData) context.Context {
	return context.WithValue(context.Background(), "rmd", rmd)
}

func RunInEntTransaction(entClient *entgen.Client, ctx context.Context, EntFunc func(tx *entgen.Tx) error) error {
	tx, err := entClient.Tx(ctx)
	if err != nil {
		return errorx.WrapENTError("ent_transaction", err)
	}

	err = EntFunc(tx)

	if r := recover(); r != nil {
		tx.Rollback()
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
