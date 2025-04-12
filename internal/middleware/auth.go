package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/TanguiLouedec/flickhive_back/internal/services"
	"github.com/TanguiLouedec/flickhive_back/pkg/logger"
	"go.uber.org/zap"
)

type contextKey string

const userIDKey = contextKey("userID")

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logger.GetLogger()

		tokenStr := r.Header.Get("Authorization")
		if tokenStr == "" {
			log.Warn("Missing authorization bearer")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(tokenStr, "Bearer ") {
			log.Warn("Invalid authorization format", zap.String("header", tokenStr))
		}

		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
		userID, err := services.ValidateJWT(tokenStr)
		if err != nil {
			log.Warn("JWT validation failed", zap.Error(err))
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		log.Info("Authenticated user", zap.String("userID", userID))
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(userIDKey).(string)
	return userID, ok
}
