package main

import (
	"sayamphoo/microservice/app"
	"sayamphoo/microservice/config"

	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load(".env")
	if err != nil {
		panic("Environment Error")
	}
	config.PingElasticsearchClient()
	// logFile := security.FileLog()
	// defer logFile.Close()
	// log.SetOutput(logFile)
	app.InitializeController()
}
