package main

import "fmt"

var a int = 0

func main() {

	fmt.Println("global variable, a = ", a)
	a = 3
	fmt.Println("global variable, a = ", a)
	a = 10
	fmt.Println("local variable, a = ", a)
	a--
	fmt.Println("local variable, a = ", a)

	fmt.Println("global variable, a = ", a)
}
