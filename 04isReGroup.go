package interview

import "strings"

func isReGroup(s1, s2 string) bool {
	len1, len2 := len(s1), len(s2)

	if len1 > 5000 || len2 > 5000 || len1 != len2 {
		return false
	}
	for _, ch := range s1 {
		if strings.Count(s1, string(ch)) != strings.Count(s2, string(ch)) {
			return false
		}
	}
	return true
}
