package controller

import (
	"awsCloud/http/services"

	"github.com/gin-gonic/gin"
)

func CraeteBucket(ctx *gin.Context) {
	res, status := services.CreateBucket(ctx)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

// Upload File controller
// @Summary Varify token & upload a new file.
// @Description You need to signedIn and give a Token in headers then "Upload Item" will execute.
// @Tags Upload Item
// @Accept  json
// @Produce  json
// @Router /upload [post]
func UploadItem(ctx *gin.Context) {
	res, status := services.UploadItem(ctx)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}

// Items list controller
// @Summary Varify token & display the items list.
// @Description You need to signedIn and give a Token in headers then "Items List" will execute.
// @Tags Items List
// @Accept  json
// @Produce  json
// @Router /items [get]
func Items(ctx *gin.Context) {
	res, status := services.GetAllItem(ctx)
	if !res.Success {
		ctx.JSON(status, res)
		return
	}

	ctx.JSON(status, res)
}
