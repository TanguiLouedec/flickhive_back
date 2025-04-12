package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/TanguiLouedec/flickhive_back/internal/middleware"
	"github.com/TanguiLouedec/flickhive_back/internal/storage"
	"github.com/TanguiLouedec/flickhive_back/pkg/logger"
	"go.uber.org/zap"
)

func GetProfile(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()

		userID, ok := middleware.GetUserIDFromContext(r.Context())
		if !ok {
			log.Warn("/me access failed: userID not in context")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		log.Info("Fetching profile for user", zap.String("userID", userID))

		user, err := storage.GetUserByID(db, userID)
		if err != nil {
			log.Error("DB error fetching user", zap.Error(err))
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		log.Info("Successfully fetched user profile", zap.String("email", user.Email))

		json.NewEncoder(w).Encode(map[string]interface{}{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
		})
	}
}
