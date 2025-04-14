package storage

import (
	"database/sql"

	"github.com/TanguiLouedec/flickhive_back/internal/models"
)

func CreateUser(db *sql.DB, user *models.User) error {
	query := `
    INSERT INTO users (id, username, email, password, created_at)
    VALUES ($1, $2, $3, $4, $5)
  `
	_, err := db.Exec(query, user.ID, user.Username, user.Email, user.Password, user.CreatedAt)
	return err
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, created_at FROM users WHERE email = $1`
	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByID(db *sql.DB, id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, created_at FROM users WHERE id = $1`
	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, created_at FROM users WHERE username = $1`
	err := db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UpdateUser(db *sql.DB, user *models.User) error {
	query := `
    UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4
  `
	_, err := db.Exec(query, user.Username, user.Email, user.Password, user.ID)
	return err
}
