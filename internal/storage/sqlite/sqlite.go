package sqlite

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(connString string) (*Storage, error) {
	const op = "storage.postgresql.New"

	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = conn.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS url (
			id SERIAL PRIMARY KEY,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: conn}, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.postgres.SaveURL"

	var id int64
	err := s.db.QueryRow(
		context.Background(),
		"INSERT INTO url(url, alias) VALUES($1, $2) RETURNING id",
		urlToSave, alias,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.postgres.GetURL"

	var url string

	err := s.db.QueryRow(context.Background(),
		"SELECT url FROM url WHERE alias=$1",
		alias,
	).Scan(&url)

	if err != nil {
		if err == pgx.ErrNoRows {
			return "", fmt.Errorf("%s: alias '%s' not found", op, alias)
		}
		return "", fmt.Errorf("%s: failed to execute query: %w", op, err)
	}

	return url, nil
}

func (s *Storage) DeleteURL(alias string) error {
	const op = "storage.postgres.delete"

	err := s.db.QueryRow(context.Background(),
		"DELETE FROM url WHERE alias=$1",
		alias)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
