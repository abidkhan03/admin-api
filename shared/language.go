package shared

import (
	"regexp"
)

var (
	// compile regular expression to remove non-alphabetic characters
	cleanUpRegex = regexp.MustCompile("[^a-zA-ZñÑüÜáéíóú ]+")
)

// CleanUpText removes non-alphabetic characters from the text
func CleanUpText(text string) string {
	return cleanUpRegex.ReplaceAllString(text, "")
}
