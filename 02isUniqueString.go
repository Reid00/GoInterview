package interview

import "strings"

func isUniqueString(s string) bool {

	if len(s) > 3000 {
		return false
	}

	// 因为不允许使用额外的存储空间，不能使用map
	// 1. 使用strings.Count
	for _, v := range s {
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true

	// 2. 使用strings.Index
	// for i, v := range s {
	// 	if strings.Index(s, string(v)) != i {
	// 		return false
	// 	}
	// }
	// return true
}
