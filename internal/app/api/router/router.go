package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"media/backend/internal/domain/users"
	"media/backend/internal/infra/db"
)

func NewRouter(pool *pgxpool.Pool) chi.Router {
	q := db.NewQueries(pool)
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: false,
	}))

	// User domain setup
	userRepo := users.NewRepository(q)
	//userSvc := user.NewService(userRepo)
	//userRoutes := user.Routes(userSvc)
}
