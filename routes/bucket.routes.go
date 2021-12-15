package routes

import (
	ctrl "awsCloud/controller"
	"awsCloud/middleware"

	"github.com/gin-gonic/gin"
)

func bucketRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth", middleware.VerifyUser())

	auth.POST("create-bucket", ctrl.CraeteBucket)
	auth.POST("upload", ctrl.UploadItem)
	auth.GET("items", ctrl.Items)
}
