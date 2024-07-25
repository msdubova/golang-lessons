package storage

import (
	"context"
	"fmt"
	"web21/internal/core"
)

func (s *Storage) GetShortURLByID(ctx context.Context, id string) (*core.ShortURL, error) {
	const q = `
		SELECT id, original_url
		FROM short_urls
		WHERE id = $1
	`

	rows, err := s.db.QueryContext(ctx, q, id)
	if err != nil {
		return nil, fmt.Errorf("selecting short URL: %w", err)
	}

	if !rows.Next() {
		return nil, fmt.Errorf("%w: %v", core.ErrURLNotFound, id)
	}

	var u core.ShortURL

	if err := rows.Scan(&u.ID, &u.OriginalURL); err != nil {
		return nil, fmt.Errorf("scanning short url: %w", err)
	}

	return &u, nil
}

func (s *Storage) CreateShortURL(ctx context.Context, shortURL core.ShortURL) error {
	const q = `
		INSERT INTO short_urls(id, original_url, user_id)
		VALUES($1, $2, $3)
	`

	_, err := s.db.ExecContext(ctx, q, shortURL.ID, shortURL.OriginalURL, shortURL.UserID)
	if err != nil {
		return fmt.Errorf("inserting short URL: %w", err)
	}

	return nil
}

func (s *Storage) CreateUserAndShortURL(ctx context.Context, user core.User, shortURL core.ShortURL) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("starting tx: %w", err)
	}

	_, err = tx.Exec("INSERT INTO users(id, nickname, password) VALUES ($1, $2, $3)",
		user.ID, user.Nickname, user.Password)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("inserting user: %w", err)
	}

	_, err = tx.Exec("INSERT INTO short_url(id, original_url, user_id) VALUES ($1, $2, $3)",
		shortURL.ID, shortURL.OriginalURL, shortURL.UserID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("inserting short URL: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commiting tx: %w", err)
	}

	return nil
}
