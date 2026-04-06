package main

import (
	"context"
	"media/backend/configs"
	"media/backend/internal/infra/db"
	"media/backend/internal/pkg/logger"
	"net/http"
)

func main() {
	cfg := configs.Load()
	ctx := context.Background()

	pool, err := db.NewPostgresPool(ctx, cfg)
	if err != nil {
		logger.Fatal("Failed to connect to database: %v", err)
	}
	defer pool.Close()
	//r := router.NewRouter(pool)
	if err := db.RunMigrations(ctx, pool); err != nil {
		logger.Fatal("Failed to run migrations: %v", err)
	}
	logger.Info("Server started on port %s", cfg.ServerPort)
	http.ListenAndServe(cfg.ServerPort, nil)

}
