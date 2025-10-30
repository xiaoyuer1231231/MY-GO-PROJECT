// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。
package main

import "fmt"

func main() {
	strs := [...]string{"flow", "flower", "flight"}
	ss := large(strs[:])
	fmt.Println(ss)
}

func large(str []string) string {

	if len(str) == 0 {
		return ""
	}
	str1 := str[0]
	count := len(str)
	for i := 1; i < count; i++ {
		str1 = lcp(str1, str[i])
		if len(str1) == 0 {
			break
		}
	}
	return str1
}
func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length {
		if str1[index] != str2[index] {
			break
		}
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
