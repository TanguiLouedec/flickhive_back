package storage

import (
	"github.com/TanguiLouedec/flickhive_back/pkg/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"database/sql"
	"fmt"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	log := logger.GetLogger()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Error("DB_URL is not set")
		return nil, fmt.Errorf("DB_URL is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Error("sql.Open error", zap.Error(err))
		return nil, fmt.Errorf("sql.Open error: %w", err)
	}

	db.SetMaxOpenConns(25)

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("db.Ping error: %w", err)
	}

	log.Info("DB successfuly pinged")

	return db, nil
}
