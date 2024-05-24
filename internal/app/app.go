package app

import (
	"table_top/internal/configs"
	"table_top/internal/dbs"
	"table_top/internal/handlers"
	"table_top/internal/migrations"
	"table_top/internal/repositories"
	"table_top/internal/routers"
	"table_top/internal/services"

	"log"

	"github.com/gin-gonic/gin"
)

func InitApp(env string) *gin.Engine {
	cfg := configs.LoadConfig(env)
	database, err := dbs.Connect(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err := migrations.Migrate(database); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	repos := repositories.InitRepositories(database)
	services := services.InitServices(repos)
	handlers.InitHandlers(services)

	router := routers.InitRouter()

	return router
}
