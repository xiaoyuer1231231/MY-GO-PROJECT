// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

// 有效字符串需满足：

//     左括号必须用相同类型的右括号闭合。
//     左括号必须以正确的顺序闭合。
//     每个右括号都有一个对应的相同类型的左括号。

package main

import "fmt"

func aa(a string) bool {
	n := len(a)
	if n%2 == 1 {
		return false
	}
	mapNew := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if mapNew[a[i]] > 0 {
			stack = append(stack, a[i])

		} else {
			fmt.Println("stack[len(stack)-1]", stack[len(stack)-1])
			if len(stack) == 0 || stack[len(stack)-1] != mapNew[a[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
func main() {
	var message string = "([{[{(}}]}])"
	ss := aa(message)
	fmt.Println(ss)
}

// func isValid(s string) bool {
// 	n := len(s)
// 	if n%2 == 1 {
// 		return false
// 	}
// 	pairs := map[byte]byte{
// 		')': '(',
// 		']': '[',
// 		'}': '{',
// 	}
// 	stack := []byte{}
// 	for i := 0; i < n; i++ {
// 		fmt.Println("pairs[s[i]]", pairs[s[i]])
// 		fmt.Println("s[i]", s[i])
// 		if pairs[s[i]] > 0 {

// 			fmt.Println("stack[len(stack)-1]", stack[len(stack)-1])
// 			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
// 				return false
// 			}
// 			stack = stack[:len(stack)-1]
// 			fmt.Println("end]", stack)

// 		} else {
// 			fmt.Println("s1", s[i])
// 			stack = append(stack, s[i])
// 		}
// 	}
// 	return len(stack) == 0
// }
