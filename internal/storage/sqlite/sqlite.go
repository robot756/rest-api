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

//func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
//	const op = "storage.sqlite.SaveURL"
//
//	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES(?, ?)")
//	if err != nil {
//		return 0, fmt.Errorf("%s: %w", op, err)
//	}
//
//	res, err := stmt.Exec(urlToSave, alias)
//	if err != nil {
//		return 0, fmt.Errorf("%s: %w", op, err, "alias already exist")
//	}
//
//	id, err := res.LastInsertId()
//	if err != nil {
//		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
//	}
//
//	return id, nil
//}

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
