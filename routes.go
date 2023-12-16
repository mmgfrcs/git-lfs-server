package main

import (
	"git-lfs-server/controllers"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitAPIRoutes(group *echo.Group) {
	ctrl := controllers.APIController{}
	group.GET("/health", ctrl.Ping)
	group.PUT("/upload/lfs/:org/:project/:oid", ctrl.UploadLocal, echojwt.JWT([]byte("")))
}

func InitLFSRoutes(group *echo.Group) {
	ctrl := controllers.LFSController{}
	group.POST("/objects/batch", ctrl.Batch)
}
