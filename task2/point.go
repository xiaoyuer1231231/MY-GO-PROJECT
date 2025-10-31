// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。

//     考察点 ：指针的使用、值传递与引用传递的区别

package main

import "fmt"

func ss(s *int) {
	ss := *s + 10
	fmt.Println(ss)
}
func spile(ptr *[]int) {
	slice := *ptr
	for v := range slice {
		fmt.Println(v)
	}
}

func main() {
	// var i int = 10
	// ss(&i)
	aa := []int{1, 2, 3, 4}
	fmt.Println("原始切片:", aa)
	spile(&aa)
	fmt.Println("修改后切片:", spile)
}
