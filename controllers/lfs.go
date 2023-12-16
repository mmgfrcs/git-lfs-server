package controllers

import (
	"git-lfs-server/models"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
)

type LFSController struct{}

func (c LFSController) Batch(ctx echo.Context) error {
	var req models.LFSBatchRequest

	ctx.Response().Header().Set(echo.HeaderContentType, "application/vnd.git-lfs+json")

	if err := ctx.Bind(&req); err != nil {
		return err
	}
	if err := ctx.Validate(&req); err != nil {
		return err
	}

	if !lo.Contains(req.Transfers, "basic") {
		return echo.NewHTTPError(409, "Unsupported transfer adapter. Supported transfer adapter: basic")
	}

	res := models.LFSBatchResponse{Transfer: "basic", HashAlgo: "sha256", Objects: make([]models.LFSBatchObjectResponse, 0)}

	if req.Operation == models.LFSOpUpload {
		for _, obj := range req.Objects {
			res.Objects = append(res.Objects, models.LFSBatchObjectResponse{OID: obj.OID, Size: obj.Size})
		}
	} else {
		return ctx.JSON(404, models.LFSBatchErrorResponse{Message: "All objects not found"})
	}

	return ctx.JSON(200, res)
}
