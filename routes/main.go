package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	var API_PREFIX = os.Getenv("API_PREFIX")
	var API_VERSION = os.Getenv("API_VERSION")

	prefix := router.Group("/" + API_PREFIX + "/" + API_VERSION)
	// user api routes
	userRoutes(prefix)
	bucketRoutes(prefix)

	return router
}
