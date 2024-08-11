package application

import (
	"time"

	"github.com/golanglowell/quick-link/internal/domain"
	"github.com/teris-io/shortid"
)

type ShortenURL struct {
	urlRepo domain.URLRepository
}

func NewShortenURL(urlRepo domain.URLRepository) *ShortenURL {
	return &ShortenURL{urlRepo: urlRepo}
}

func (uc *ShortenURL) Execute(longURL string) (*domain.URL, error) {
	shortCode, _ := shortid.Generate()
	link := &domain.URL{
		LongURL:   longURL,
		ShortCode: shortCode,
		CreatedAt: time.Now(),
	}
	err := uc.urlRepo.Save(link)
	if err != nil {
		return nil, err
	}
	return link, nil
}
