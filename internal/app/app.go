package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"table_top/internal/configs"
	"table_top/internal/dbs"
	"table_top/internal/handlers"
	"table_top/internal/migrations"
	"table_top/internal/repositories"
	"table_top/internal/routers"
	"table_top/internal/services"
)

func Run(env string) {
	// Создание основного контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Настройка для перехвата сигналов завершения
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		<-sig
		cancel()
	}()

	cfg := configs.LoadConfig(env)
	database, err := dbs.Connect(cfg)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err := migrations.Migrate(database); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	repos := repositories.InitRepositories(database)
	svc := services.InitServices(repos)
	handlers.InitHandlers(svc)

	router := routers.InitRouter(ctx)

	log.Println("Application started successfully on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
