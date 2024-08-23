package utils

import (
	"regexp"
	"strings"
)

// GenerateSlug converts a title into a URL-friendly slug
func GenerateSlug(title string) string {
	// Convert to lowercase
	title = strings.ToLower(title)
	// Replace non-alphanumeric characters with hyphens
	re := regexp.MustCompile(`[^a-z0-9]+`)
	slug := re.ReplaceAllString(title, "-")
	// Trim hyphens from the start and end
	return strings.Trim(slug, "-")
}
