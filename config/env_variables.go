package config

import "os"

var PORT = os.Getenv("PORT")
var API_PREFIX = os.Getenv("API_PREFIX")
var API_VERSION = os.Getenv("API_VERSION")
var AWS_REGION = os.Getenv("AWS_REGION")
var TABLE_NAME = os.Getenv("TABLE_NAME")
var ACCESS_SECRET = os.Getenv("ACCESS_SECRET")
var REFRESH_SECRET = os.Getenv("REFRESH_SECRET")
