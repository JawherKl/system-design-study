package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/JawherKl/url-shortener/internal/model"
	"os"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo() *PostgresRepo {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_DSN"))
	if err != nil {
		panic(err)
	}
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) SaveURL(url model.URL) error {
	_, err := r.db.Exec("INSERT INTO urls (original, short) VALUES ($1, $2)", url.Original, url.Short)
	return err
}

func (r *PostgresRepo) GetOriginalURL(hash string) (string, error) {
	var original string
	err := r.db.QueryRow("SELECT original FROM urls WHERE short = $1", hash).Scan(&original)
	return original, err
}