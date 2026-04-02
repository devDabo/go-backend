package main

import (
	"context"
	"media/backend/configs"
	"media/backend/internal/infra/db"
	"media/backend/internal/pkg/logger"
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
	logger.Info("Server started on port %s", cfg.ServerPort)

}
