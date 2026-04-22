package users

import (
	"context"
	"github.com/jackc/pgx/v5/pgtype"
	"media/backend/internal/pkg/database/sqlc"
)

type Repository interface {
	Create(ctx context.Context, u User) (User, error)
}

type repository struct {
	queries *sqlc.Queries
}

func NewRepository(q *sqlc.Queries) Repository {
	return &repository{q}
}

func (r *repository) Create(ctx context.Context, u User) (User, error) {
	params := sqlc.CreateUserParams{
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Bio:         u.Bio,
		AvatarUrl:   u.AvatarURL,
	}

}
