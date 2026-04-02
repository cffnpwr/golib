package path

import "strings"

const maxFileNameLength = 255

// IsValidName reports whether s is a valid filename element
// (a single path component with no path separators).
// It checks cross-platform constraints common to all supported OSes.
func IsValidName(s string) bool {
	if s == "" {
		return false
	}
	if len(s) > maxFileNameLength {
		return false
	}
	if strings.Contains(s, "\x00") {
		return false
	}
	if strings.ContainsAny(s, `/\`) {
		return false
	}

	return isValidName(s)
}
