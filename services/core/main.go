package main

import (
	"log"
	"os"

	"vou/services/core/api"

	"github.com/joho/godotenv"
)

func main() {
	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}

	app := api.NewApp()
	app.InitRouter()

	port, found := os.LookupEnv("CORE_PORT")
	if !found {
		port = "8080"
	}

	app.Engine.Run(":" + port)
}
