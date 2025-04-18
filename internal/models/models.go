package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"_"`
	CreatedAt time.Time `json:"created_at"`
}

type Movie struct {
	ID        uuid.UUID `json:"id"`
	TMDBID    int       `json:"tmdb_id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}
