// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

//     考察点 ：组合的使用、方法接收者。

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person     Person
	EmployeeID int
}

func (c *Employee) PrintInfo() {
	fmt.Println("id为", c.EmployeeID)
	fmt.Println("年龄", c.Person.Age)
	fmt.Println("名称", c.Person.Name)
}

func main() {
	person := Person{Name: "张三", Age: 123}
	person1 := Employee{Person: person, EmployeeID: 1}
	person1.PrintInfo()

}
