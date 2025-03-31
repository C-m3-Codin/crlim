package internal

import (
	"net/url"
	"strings"
)

// NormalizeURL removes query parameters and fragments
func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	cleanedPath := strings.TrimSuffix(parsedURL.Host+parsedURL.Path, "/")
	return cleanedPath, nil
}
