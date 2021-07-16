package interview

import (
	"strings"
	"unicode"
)

func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}

	for _, char := range s {
		if char != ' ' && !unicode.IsLetter(char) {
			return s, false
		}
	}
	return strings.Replace(s, " ", "%20", -1), true
}
