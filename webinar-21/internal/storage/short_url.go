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
		INSERT INTO short_urls(id, original_url)
		VALUES($1, $2)
	`

	_, err := s.db.ExecContext(ctx, q, shortURL.ID, shortURL.OriginalURL)
	if err != nil {
		return fmt.Errorf("inserting short URL: %w", err)
	}

	return nil
}
