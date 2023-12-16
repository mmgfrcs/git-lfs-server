package controllers

import "github.com/labstack/echo/v4"

type APIController struct{}

func (c APIController) Ping(ctx echo.Context) error {
	return ctx.NoContent(204)
}

func (c APIController) UploadLocal(ctx echo.Context) error {
	//ctx.Request().Body
	return nil
}
