package main

import "fmt"

// 普通函数
func add(a, b int) int {
	return a + b
}

// 多返回值函数
func swap(x int, y string) (string, int) {
	return y, x
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}
func main() {
	area := Rectangle{width: 10, height: 5}
	// result := add(3, 5)          // 直接调用
	fmt.Println("3+5 =", area.area()) // 输出: 3+5 = 8

	// a, b := swap(1, "world")
	// fmt.Println(a, b) // 输出: world hello
}
