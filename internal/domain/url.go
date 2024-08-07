package domain

import "time"

type URL struct {
	ID        string
	LongURL   string
	ShortCode string
	CreatedAt time.Time
	Clicks    int
}

type URLRepository interface {
	SaveURL(url *URL) error
	FindByShortCode(shortCode string) (*URL, error)
	IncrementClicks(shortCode string) error
}

type ShortCodeGenerator interface {
	Generate() string
}
