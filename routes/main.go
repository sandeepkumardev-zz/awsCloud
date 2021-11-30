package routes

import (
	"os"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	// Api Prefix
	apiPrefix := os.Getenv("API_PREFIX")
	// version 1
	apiVersion := os.Getenv("API_VERSION")

	prefix := router.Group("/" + apiPrefix + "/" + apiVersion)
	// user api routes
	userRoutes(prefix)

	return router
}
