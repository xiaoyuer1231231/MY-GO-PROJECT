package main

import "fmt"

func main() {
	a := "test string"

	// 1. 基本用法
	switch a {
	case "test":
		fmt.Println("a = ", a)
	case "s":
		fmt.Println("a = ", a)
	case "t", "test string": // 可以匹配多个值，只要一个满足条件即可
		fmt.Println("catch in a test, a = ", a)
	case "n":
		fmt.Println("a = not")
	default:
		fmt.Println("default case")
	}

	// 变量b仅在当前switch代码块内有效
	switch b := 5; b {
	case 1:
		fmt.Println("b = 1")
	case 2:
		fmt.Println("b = 2")
	case 3, 4:
		fmt.Println("b = 3 or 4")
	case 5:
		fmt.Println("b = 5")
	default:
		fmt.Println("b = ", b)
	}

	// 不指定判断变量，直接在case中添加判定条件
	b := 5
	switch {
	case a == "t":
		fmt.Println("a = t")
	case b == 3:
		fmt.Println("b = 5")
	case b == 5, a == "test string":
		fmt.Println("a = test string; or b = 5")
	default:
		fmt.Println("default case")
	}

}
