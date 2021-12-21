package routes

import (
	"os"

	docs "awsCloud/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	var API_PREFIX = os.Getenv("API_PREFIX")
	var API_VERSION = os.Getenv("API_VERSION")

	docs.SwaggerInfo.BasePath = "/" + API_PREFIX + "/" + API_VERSION

	prefix := router.Group("/" + API_PREFIX + "/" + API_VERSION)
	// user api routes
	userRoutes(prefix)
	bucketRoutes(prefix)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
