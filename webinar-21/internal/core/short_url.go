package core

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/google/uuid"
)

type ShortURL struct {
	ID          string
	OriginalURL string
}

type ShortURLStorage interface {
	GetShortURLByID(ctx context.Context, id string) (*ShortURL, error)
	CreateShortURL(ctx context.Context, shortURL ShortURL) error
}

type ShortURLService struct {
	Storage ShortURLStorage
}

func (s *ShortURLService) OpenShortURL(ctx context.Context, id string) (*ShortURL, error) {
	shortURL, err := s.Storage.GetShortURLByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting short URL: %w", err)
	}

	return shortURL, nil
}

func (s *ShortURLService) GetShortURLByID(ctx context.Context, id string) (*ShortURL, error) {
	shortURL, err := s.Storage.GetShortURLByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting short URL: %w", err)
	}

	return shortURL, nil
}

var (
	ErrInvalidURL  = errors.New("invalid URL")
	ErrURLNotFound = errors.New("URL not found")
)

func (s *ShortURLService) CreateShortURL(ctx context.Context, rawURL string) (*ShortURL, error) {
	if _, err := url.Parse(rawURL); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidURL, rawURL)
	}

	shortURL := ShortURL{
		ID:          uuid.NewString(),
		OriginalURL: rawURL,
	}

	err := s.Storage.CreateShortURL(ctx, shortURL)
	if err != nil {
		return nil, fmt.Errorf("creating short URL: %w", err)
	}

	return &shortURL, nil
}
