package domain

import "errors"

var (
	ErrURLNotFound        = errors.New("URL not found")
	ErrURLInvalidURL      = errors.New("invalid URL")
	ErrDuplicateShortCode = errors.New("duplicate short code")
	ErrInternalServer     = errors.New("internal server error")
)
