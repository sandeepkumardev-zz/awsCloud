package main

import (
	"awsCloud/database/config"
	"awsCloud/http/routes"
	"fmt"

	"github.com/joho/godotenv"
)

// @title AWS Cloud API Documentation.
// @version 1.0.0
// @description A service where users can register and store there files.
// @termsOfService http://swagger.io/terms/

// @contact.name Sandeep kumar
// @contact.email sandeepypb@gmail.com

// @host localhost:3000
// @BasePath /api/v1

func main() {
	fmt.Println("AWS cloud storage ...")

	// load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading environment variables.")
	}

	// stabling the database
	_, dbErr := config.ConnectionDB()
	if dbErr != nil {
		panic("Error in stabling database connection: " + dbErr.Error())
	}

	// start server
	var port = config.GetEnvVar().PORT
	router := routes.RouterSetup()
	router.Run(":" + port)
}
