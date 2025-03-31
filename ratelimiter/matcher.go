package ratelimiter

import (
	"fmt"
	"regexp"
	"strings"
)

// WildcardToRegex converts wildcard patterns (*, {param}) to regex
func WildcardToRegex(pattern string) string {
	pattern = strings.ReplaceAll(pattern, ".", "\\.")                            // Escape dots
	pattern = strings.ReplaceAll(pattern, "*", "[^/]+")                          // Convert * to regex
	pattern = regexp.MustCompile(`\{[^}]+\}`).ReplaceAllString(pattern, `[^/]+`) // Convert {param} to regex
	return "^" + pattern + "$"
}

// MatchWithWildcard checks if a given URL matches a wildcard pattern
func MatchWithWildcard(url string, patterns map[string]*TokenBucket) (string, bool) {
	for pattern := range patterns {
		regexPattern := WildcardToRegex(pattern)
		matched, _ := regexp.MatchString(regexPattern, url)
		if matched {
			fmt.Println(" match", pattern, url)
			return pattern, true
		} else {
			fmt.Println("didnt match", pattern, url)
		}
	}
	return "", false
}
