package sqlite

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := pgx.Connect(context.Background(), storagePath)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = db.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS url (
	    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	    alias TEXT NOT NULL UNIQUE,
	    url TEXT NOT NULL);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
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

//func (s *Storage) GetURL(alias string) (string, error) {
//	const op = "storage.sqlite.GetURL"
//
//	stmt, err := s.db.Prepare("SELECT url FROM url WHERE alias=?")
//	if err != nil {
//		return "", fmt.Errorf("failed to prepate stmt %s: %w", err)
//	}
//
//	var url string
//	err = stmt.QueryRow(alias).Scan(&url)
//	if err != nil {
//		return "", fmt.Errorf("%s: alias %s not found", op, alias)
//	}
//
//	return url, nil
//}
//
//func (s *Storage) DeleteURL(alias string) error {
//	const op = "storage.sqllite.delete"
//
//	stmt, err := s.db.Prepare("DELETE FROM url WHERE alias=?")
//	if err != nil {
//		return fmt.Errorf("%s: %w", op, err)
//	}
//
//	res, err := stmt.Exec(alias)
//	if err != nil {
//		return fmt.Errorf("%s: %w", op, err)
//	}
//
//	rowsAffected, err := res.RowsAffected()
//	if err != nil {
//		return fmt.Errorf("%s: rows is not found", op, alias)
//	}
//
//	if rowsAffected == 0 {
//		return fmt.Errorf("%s: alias %s: not found", op, alias)
//	}
//
//	return nil
//}
