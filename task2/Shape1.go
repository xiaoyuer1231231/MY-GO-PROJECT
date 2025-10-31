// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。

// 考察点 ：接口的定义与实现、面向对象编程风格。
package main

import (
	"fmt"
)

type Rectangle struct {
	i int
}
type Circle struct {
	i int
}
type Shape interface {
	Area(int)
	Perimeter(int)
}

func (Area *Rectangle) Area(i int) {
	fmt.Println("打印Area", i)
}
func (Perimeter *Circle) Perimeter(i int) {
	fmt.Println("Perimeter", i)
}
func main() {
	c := &Rectangle{i: 1}
	s := &Circle{i: 1}
	c.Area(2)
	s.Perimeter(3)
}
