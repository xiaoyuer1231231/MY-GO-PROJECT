package main

import "fmt"

func main() {
	// 切片
	s := make([]string, 10)
	s[0] = "Hello"
	s[1] = "World"

	// 数组
	var a [10]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println("=== 初始状态 ===")
	fmt.Printf("切片 s: 长度=%d, 容量=%d, 值=%v\n", len(s), cap(s), s)
	fmt.Printf("数组 a: 长度=%d, 值=%v\n", len(a), a)

	fmt.Println("\n=== 动态扩容能力 ===")
	// 切片可以动态追加
	s = append(s, "New Element")
	fmt.Printf("切片追加后: 长度=%d, 容量=%d\n", len(s), cap(s))

	// 数组不能追加（编译错误）
	// a = append(a, "New Element")  // 错误！

	fmt.Println("\n=== 赋值行为 ===")
	// 切片赋值是引用传递
	s2 := s
	s2[0] = "Modified"
	fmt.Printf("原切片 s[0] = %s\n", s[0])   // Modified（被修改）
	fmt.Printf("新切片 s2[0] = %s\n", s2[0]) // Modified

	// 数组赋值是值传递（创建副本）
	a2 := a
	a2[0] = "Modified"
	fmt.Printf("原数组 a[0] = %s\n", a[0])   // Hello（保持不变）
	fmt.Printf("新数组 a2[0] = %s\n", a2[0]) // Modified

}
