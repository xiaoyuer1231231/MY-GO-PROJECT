package main

import "fmt"

// func main() {
// 	// 三维数组
// 	a := [3][2][2]int{
// 		{{0, 1}, {2, 3}},
// 		{{4, 5}, {6, 7}},
// 		{{8, 9}, {10, 11}},
// 	}

// 	layer1 := a[0]
// 	layer2 := a[0][1]
// 	element := a[0][1][1]
// 	fmt.Println(layer1)
// 	fmt.Println(layer2)
// 	fmt.Println(element)

// 	// 多维数组遍历时，需要使用嵌套for循环遍历
// 	for i, v := range a {
// 		fmt.Println("index = ", i, "value = ", v)
// 		for j, inner := range v {
// 			fmt.Println("inner, index = ", j, "value = ", inner)
// 			for k, inner := range inner {
// 				fmt.Println("inner, index = ", k, "value = ", inner)
// 			}
// 		}
// 	}
// }
type Custom struct {
	i int
}

var carr [5]*Custom = [5]*Custom{
	{6},
	{7},
	{8},
	{9},
	{10},
}

func main() {
	a := [5]int{5, 4, 3, 2, 1}
	fmt.Println("before all, a = ", a)
	for i := range carr {
		fmt.Printf("in main func, carr[%d] = %p, value = %v \n", i, &carr[i], *carr[i])
	}
	printFuncParamPointer(carr)

	receiveArray(a)
	fmt.Println("after receiveArray, a = ", a)

	receiveArrayPointer(&a)
	fmt.Println("after receiveArrayPointer, a = ", a)
}

func receiveArray(param [5]int) {
	fmt.Println("in receiveArray func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArray func, after modify, param = ", param)
}

func receiveArrayPointer(param *[5]int) {
	fmt.Println("in receiveArrayPointer func, before modify, param = ", param)
	param[1] = -5
	fmt.Println("in receiveArrayPointer func, after modify, param = ", param)
}

func printFuncParamPointer(param [5]*Custom) {
	for i := range param {
		fmt.Printf("in printFuncParamPointer func, param[%d] = %p, value = %v \n", i, &param[i], *param[i])
	}
}
