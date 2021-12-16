package config

import (
	"awsCloud/database/models"
	"os"
)

type Config models.Config

func GetEnvVar() Config {
	PORT := os.Getenv("PORT")
	API_PREFIX := os.Getenv("API_PREFIX")
	API_VERSION := os.Getenv("API_VERSION")

	AWS_REGION := os.Getenv("AWS_REGION")
	AWS_SECRET := os.Getenv("AWS_SECRET")
	AWS_SECRET_KEY := os.Getenv("AWS_SECRET_KEY")
	TABLE_NAME := os.Getenv("TABLE_NAME")

	ACCESS_SECRET := os.Getenv("ACCESS_SECRET")
	REFRESH_SECRET := os.Getenv("REFRESH_SECRET")

	return Config{PORT, API_PREFIX, API_VERSION, TABLE_NAME, ACCESS_SECRET, REFRESH_SECRET, AWS_REGION, AWS_SECRET, AWS_SECRET_KEY}
}
