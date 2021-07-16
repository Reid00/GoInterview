package interview

func reversedString(s string) (string, bool) {

	str := []rune(s)
	// 先判断长度
	// 1. 转化成[]rune 类型
	if len(str) > 5000 {
		return s, false
	}
	// 2. 用utf8 包
	// if utf8.RuneCountInString(s) > 5000 {
	// 	return s, false
	// }

	// 在字符串中间，将s[0] 与 s[len-1] 进行替换
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}
	return string(str), true
}
