package memory

import "github.com/golanglowell/quick-link/internal/domain"

type URLRepository struct {
	urls map[string]domain.URL
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		urls: make(map[string]domain.URL),
	}
}

func (u *URLRepository) Save(url domain.URL) error {
	if err := url.Validate(); err != nil {
		return err
	}
	if _, exists := u.urls[url.ShortCode]; exists {
		return domain.ErrDuplicateShortCode
	}

	u.urls[url.ShortCode] = url

	return nil
}

func (u *URLRepository) FindByShortCode(shortCode string) (*domain.URL, error) {
	url, exists := u.urls[shortCode]
	if !exists {
		return nil, domain.ErrURLNotFound
	}

	return &url, nil
}

func (u *URLRepository) IncrementClicks(shortCode string) error {
	url, exists := u.urls[shortCode]
	if !exists {
		return domain.ErrURLNotFound
	}

	url.Clicks++

	return nil
}
