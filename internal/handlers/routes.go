package handlers

import (
	"database/sql"
	"net/http"

	"github.com/TanguiLouedec/flickhive_back/internal/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewRouter(log *zap.Logger, db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// Public routes
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy"))
	})

	r.Post("/signup", SignUp(db))
	r.Post("/login", Login(db))

	// Auth routes
	r.Group(func(r chi.Router) {
		r.Use(middleware.JWTAuthMiddleware)
		r.Get("/me", GetProfile(db))
	})

	return r
}
