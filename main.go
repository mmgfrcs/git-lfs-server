package main

import (
	"git-lfs-server/config"
	"git-lfs-server/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conf := config.LoadConfig()
	app := echo.New()
	app.Logger.SetLevel(2)
	app.HideBanner = true
	app.Binder = new(models.LFSBinder)
	app.Validator = models.NewValidator()

	app.Use(middleware.Logger())
	app.Use(middleware.RequestID())

	InitAPIRoutes(app.Group("/api"))
	InitLFSRoutes(app.Group("/:org/:proj"))

	app.Logger.Info("Starting LFS server")
	app.Logger.Fatal(app.Start(":" + conf.Port()))
}
