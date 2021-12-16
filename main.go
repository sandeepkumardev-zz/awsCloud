package main

import (
	"awsCloud/database/config"
	"awsCloud/http/routes"
	"fmt"

	"github.com/joho/godotenv"
)

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
