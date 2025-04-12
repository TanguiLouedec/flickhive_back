package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/TanguiLouedec/flickhive_back/internal/models"
	"github.com/TanguiLouedec/flickhive_back/internal/services"
	"github.com/TanguiLouedec/flickhive_back/internal/storage"
	"github.com/google/uuid"
)

func SignUp(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		hashedPwd, err := services.HashPassword(input.Password)
		if err != nil {
			http.Error(w, "Password hash error", http.StatusInternalServerError)
			return
		}

		user := &models.User{
			ID:        uuid.New(),
			Username:  input.Username,
			Email:     input.Email,
			Password:  hashedPwd,
			CreatedAt: time.Now(),
		}

		err = storage.CreateUser(db, user)
		if err != nil {
			http.Error(w, "User creation failed", http.StatusInternalServerError)
			return
		}

		token, err := services.GenerateJWT(user.ID.String())
		if err != nil {
			http.Error(w, "Error generating JWT token", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		user, err := storage.GetUserByEmail(db, input.Email)
		if err != nil || user == nil {
			http.Error(w, "User  not found", http.StatusUnauthorized)
			return
		}

		if !services.CheckPasswordHash(input.Password, user.Password) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}

		token, err := services.GenerateJWT(user.ID.String())
		if err != nil {
			http.Error(w, "Error generation JWT token", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"token": token})
	}
}
