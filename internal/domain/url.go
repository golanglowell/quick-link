package domain

import (
	"fmt"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type URL struct {
	ID        string    `validate:"required"`
	LongURL   string    `validate:"required"`
	ShortCode string    `validate:"required"`
	CreatedAt time.Time `validate:"required"`
	Clicks    int
}

func (u URL) Validate() error {
	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		return fmt.Errorf("url struct validation error: %w", err)
	}
	return nil
}

type URLRepository interface {
	SaveURL(url *URL) error
	FindByShortCode(shortCode string) (*URL, error)
	IncrementClicks(shortCode string) error
}

type ShortCodeGenerator interface {
	Generate() string
}
