package routes

import (
	"awsCloud/config"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	prefix := router.Group("/" + config.API_PREFIX + "/" + config.API_VERSION)
	// user api routes
	userRoutes(prefix)
	bucketRoutes(prefix)

	return router
}
