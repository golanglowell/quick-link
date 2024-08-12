package validator

import (
	"net/url"
)

func ValidateURL(urlStr string) error {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}

	return nil
}
