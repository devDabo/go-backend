package db

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"media/backend/internal/pkg/logger"
	sqlmigrations "media/backend/sql"

	"github.com/pressly/goose/v3"
)

func RunMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	const migrationsDir = "migrations"
	logger.Info("Running migrations from embedded FS: %s", migrationsDir)

	goose.SetBaseFS(sqlmigrations.FS)

	sqlDB := stdlib.OpenDBFromPool(pool)
	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			logger.Error("Failed to close DB connection: %v", err)
		}
	}(sqlDB)
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}
	err = goose.UpContext(ctx, sqlDB, migrationsDir)
	if err != nil {
		return err
	}
	return nil
}
