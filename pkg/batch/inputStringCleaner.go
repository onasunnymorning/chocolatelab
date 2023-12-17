package batch

import (
	"strings"
)

// Removes newlines(\n), tabs(\t), and carriage(\r) returns from a string and replaces them with spaces
// Removes multiple spaces and replaces them with one space
// Trims leading and trailing spaces
func CleanInputString(s string) string {
	s = RemoveNewlines(s)
	s = RemoveTabs(s)
	s = RemoveCarriageReturns(s)
	s = ReplaceMultipleSpaces(s)
	return strings.TrimSpace(s)
}

// Remove \n (newlines) from a string and replace with a space
func RemoveNewlines(s string) string {
	return strings.ReplaceAll(s, "\n", " ")
}

// Remove \t (tabs) from a string and replace with a space
func RemoveTabs(s string) string {
	return strings.ReplaceAll(s, "\t", " ")
}

// Remove \r (carriage returns) from a string and replace with a space
func RemoveCarriageReturns(s string) string {
	return strings.ReplaceAll(s, "\r", " ")
}

// Replace multiple spaces with one space
func ReplaceMultipleSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// Runs StandardizeString on all elements of a slice of strings
func NormalizeStringSlice(s []string) []string {
	for i, v := range s {
		s[i] = CleanInputString(v)
	}
	return s
}
