package stringex

import (
	"strings"
)

// Reverse will reverse a string.
func Reverse(str string) string {
	// convert string to go single characters (runes)
	runes := []rune(str)

	// while length > 0 and idx is less than jdx iterate flipper
	for idx, jdx := 0, len(runes)-1; idx < jdx; idx, jdx = idx+1, jdx-1 {
		// flipper
		runes[idx], runes[jdx] = runes[jdx], runes[idx]
	}
	return string(runes)
}

// Ellipse convert a long string to a shorter ellipsed string.
func Ellipse(str string, max int, ellipses string) string {
	// convert string to go single characters (runes)
	runes := []rune(str)
	es := len(ellipses)

	// if too long truncate with tailing ellipses
	if len(runes) > max {
		return string(runes[:max-es]) + ellipses
	}

	return str
}

func StringArrayContains(array []string, value string, caseInsensitive bool) bool {
	for _, item := range array {
		if caseInsensitive {
			if strings.EqualFold(item, value) {
				return true
			}
		} else {
			if item == value {
				return true
			}
		}
	}

	return false
}
