package handlers

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewRouter(log *zap.Logger, db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Healthy"))
	})

	r.Post("/signup", SignUp(db))
	r.Post("/login", Login(db))

	return r
}
