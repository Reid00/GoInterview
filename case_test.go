// 单元测试

package interview

import (
	"fmt"
	"testing"
)

func TestCase(t *testing.T) {
	fmt.Println("交替打印数字和字母测试:")
	PrintNumAndChar()
	fmt.Println()
	fmt.Println("====")
	fmt.Println("判断字符串中字符是否全都不同:")
	ret := isUniqueString("abcdefg")
	fmt.Println(ret)
	fmt.Println("翻转字符串:")
	res, _ := reversedString("abcdeffffg")
	fmt.Println(res)
	fmt.Println("判断两个给定的字符串排序后是否一致:")
	ok := isReGroup("判断两个给定的字符串排序后是否一致", "判断两个给是否一致定的字符串排序后")
	fmt.Println(ok)
	fmt.Println("字符串替换问题:")
	res, ok = replaceBlank("this is good")
	fmt.Println(res, ok)
	fmt.Println("两个Gorutine 生成随机数: ")
	generateInt()
}
