package db

import (
	"context"
	"fmt"
	"media/backend/configs"
	"media/backend/internal/pkg/database/sqlc"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, cfg *configs.Config) (*pgxpool.Pool, error) {
	dbURL := strings.TrimSpace(cfg.DBURL)

	if dbURL == "" {
		return nil, fmt.Errorf("DB_URL is not set")
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("unable to ping the database: %v", err)
	}

	return pool, nil
}

func NewQueries(pool *pgxpool.Pool) *sqlc.Queries {
	return sqlc.New(pool)
}
