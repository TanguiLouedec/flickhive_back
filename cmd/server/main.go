package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TanguiLouedec/flickhive_back/internal/handlers"
	"github.com/TanguiLouedec/flickhive_back/internal/storage"
	"github.com/TanguiLouedec/flickhive_back/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	logger.InitLogger()
	log := logger.GetLogger()
	defer log.Sync()

	db, err := storage.ConnectDB()
	if err != nil {
		log.Fatal("Could not connect to DB", zap.Error(err))
	}
	log.Info("DB connected", zap.String("url", os.Getenv("DB_URL")))
	defer db.Close()

	r := handlers.NewRouter(log, db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Info("🚀 Server running", zap.String("port", port))
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Server error", zap.Error(err))
	}
}
