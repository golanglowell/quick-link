package application

import (
	"fmt"
	"time"

	"github.com/golanglowell/quick-link/internal/domain"
	"github.com/golanglowell/quick-link/pkg/validator"
	"github.com/teris-io/shortid"
)

type ShortenURL struct {
	urlRepo domain.URLRepository
}

func NewShortenURL(urlRepo domain.URLRepository) *ShortenURL {
	return &ShortenURL{urlRepo: urlRepo}
}

func (a *ShortenURL) Execute(longURL string) (*domain.URL, error) {
	err := validator.ValidateURL(longURL)
	if err != nil {
		return nil, fmt.Errorf("invalid url: %w", err)
	}

	shortCode, err := shortid.Generate()
	if err != nil {
		return nil, fmt.Errorf("failed to generate short id: %w", err)
	}

	link := &domain.URL{
		LongURL:   longURL,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
	}
	err = a.urlRepo.Save(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}
