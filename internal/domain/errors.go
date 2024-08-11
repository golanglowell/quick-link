package domain

// Todo: move
import "errors"

var (
	ErrURLNotFound        = errors.New("URL not found")
	ErrURLInvalidURL      = errors.New("invalid URL")
	ErrURLInvalidInput    = errors.New("invalid input")
	ErrDuplicateShortCode = errors.New("duplicate short code")
	ErrInternalServer     = errors.New("internal server error")
)
