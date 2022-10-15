package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Configuration struct {
	ENV string

	PORT string

	ClientAppURL string

	MongoDNS string
	MongoDB  string
}

var config Configuration

func InitConfig() error {
	config.ENV = os.Getenv("ENV")

	if config.ENV == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	godotenv.Load()

	config.PORT = os.Getenv("PORT")

	config.ClientAppURL = os.Getenv("ClientAppURL")

	config.MongoDNS = os.Getenv("MongoDNS")
	config.MongoDB = os.Getenv("MongoDB")

	return nil
}
