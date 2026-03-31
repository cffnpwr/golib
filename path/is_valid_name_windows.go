//go:build windows

package path

import (
	"slices"
	"strings"
)

// windowsReservedNames lists the reserved device names on Windows.
// These names are invalid as filenames, with or without an extension.
var windowsReservedNames = []string{
	"CON", "PRN", "AUX", "NUL",
	"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
	"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
}

func isValidName(s string) bool {
	// Control characters 0x01–0x1F are forbidden.
	if strings.ContainsFunc(s, func(r rune) bool {
		return 0x01 <= r && r <= 0x1F
	}) {
		return false
	}

	// Windows-specific forbidden characters.
	if strings.ContainsAny(s, `<>:"|?*`) {
		return false
	}

	// Reserved names are invalid regardless of extension.
	// Strip the extension (everything from the first dot) for comparison.
	base := strings.ToUpper(s)
	if dot := strings.IndexByte(base, '.'); dot >= 0 {
		base = base[:dot]
	}

	if slices.Contains(windowsReservedNames, base) {
		return false
	}

	return true
}
