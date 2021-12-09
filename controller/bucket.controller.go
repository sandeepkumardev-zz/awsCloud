package controller

import (
	"awsCloud/services"

	"github.com/gin-gonic/gin"
)

func UploadItem(ctx *gin.Context) {
	res, status := services.UploadItem(ctx)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

func GetAllItem(ctx *gin.Context) {
	res, status := services.GetAllItem(ctx)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}
