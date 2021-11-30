package routes

import "github.com/gin-gonic/gin"

func userRoutes(router *gin.RouterGroup) {
	router.POST("signin")
	router.POST("signup")
	router.POST("refresh")
}
