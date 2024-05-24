package main

import (
	"log"
	"os"

	"table_top/internal/app"

	_ "table_top/docs" // Импортируйте сгенерированные Swagger-документы
)

// @title Table Top API
// @version 1.0
// @description This is a sample server for a table top game.
// @host localhost:8080
// @BasePath /

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	router := app.InitApp(env)

	log.Println("Application started successfully on :8080")
	router.Run(":8080")
}
