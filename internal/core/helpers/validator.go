package helpers

import (
	"errors"
	"strings"

	"github.com/asaskevich/govalidator"
)

func ValidateURL(url string) error {
	// check if url is a URL
	if !govalidator.IsURL(url) {
		return errors.New("not a valid url")
	}

	// remove domain error for localhost:8080
	config := LoadEnv(".")
	domain := config.Domain
	if url == domain {
		return errors.New("url is domain; url can not be domain")
	}

	invalidChars := []string{"https://", "http://", "www."}
	expectedURL := ""
	for _, char := range invalidChars {
		expectedURL = strings.Replace(url, char, "", -1)
	}

	if expectedURL == domain {
		return errors.New("url is domain; url can not be domain")
	}

	return nil
}

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}

	return url
}
