package routes

import (
	ctrl "awsCloud/controller"
	"awsCloud/middleware"

	"github.com/gin-gonic/gin"
)

func bucketRoutes(router *gin.RouterGroup) {
	router.POST("upload", middleware.VerifyUser(), ctrl.UploadItem)
	router.POST("itemlist", middleware.VerifyUser(), ctrl.GetAllItem)
}
