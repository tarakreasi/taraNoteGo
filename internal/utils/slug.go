package utils

import (
	"regexp"
	"strings"
)

// GenerateSlug creates a URL-friendly slug from a string
func GenerateSlug(input string) string {
	// Convert to lower case
	slug := strings.ToLower(input)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove all non-alphanumeric characters except hyphens
	reg := regexp.MustCompile("[^a-z0-9-]+")
	slug = reg.ReplaceAllString(slug, "")

	// Remove duplicate hyphens
	reg = regexp.MustCompile("-+")
	slug = reg.ReplaceAllString(slug, "-")

	// Trim hyphens from start and end
	slug = strings.Trim(slug, "-")

	// Fallback if empty
	if slug == "" {
		slug = "untitled"
	}

	return slug
}
