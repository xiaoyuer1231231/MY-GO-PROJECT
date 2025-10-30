package main

import "fmt"

// func main() {
// 	s := make([]int, 3, 6)
// 	fmt.Println("initial, s =", s)
// 	s[1] = 2
// 	fmt.Println("after set position 1, s =", s)

// 	s2 := append(s, 4)
// 	fmt.Println("after append, s2 length:", len(s2))
// 	fmt.Println("after append, s2 capacity:", cap(s2))
// 	fmt.Println("after append, s =", s)
// 	fmt.Println("after append, s2 =", s2)

// 	s[0] = 1024
// 	fmt.Println("after set position 0, s =", s)
// 	fmt.Println("after set position 0, s2 =", s2)

// 	appendInFunc(s)
// 	fmt.Println("after append in func, s =", s)
// 	fmt.Println("after append in func, s2 =", s2)
// }

// func appendInFunc(param []int) {
// 	param = append(param, 1022)
// 	fmt.Println("in func, param =", param)
// 	param[2] = 512
// 	fmt.Println("set position 2 in func, param =", param)
// }
func main() {
	s := make([]int, 2, 2)
	fmt.Println("initial, s =", s)

	s2 := append(s, 4)
	fmt.Println("after append, s length:", len(s))
	fmt.Println("after append, s capacity:", cap(s))

	fmt.Println("after append, s2 length:", len(s2))
	fmt.Println("after append, s2 capacity:", cap(s2))
	fmt.Println("after append, s =", s)
	fmt.Println("after append, s2 =", s2)

	s[0] = 1024
	fmt.Println("after set position 0, s =", s)
	fmt.Println("after set position 0, s2 =", s2)

	appendInFunc(s2)
	fmt.Println("after append in func, s2 =", s2)
}

func appendInFunc(param []int) {
	param1 := append(param, 511)
	param2 := append(param1, 512)
	fmt.Println("in func, param1 =", param1)
	fmt.Println("in func, param1 =", param2)
	param2[2] = 500
	fmt.Println("set position 2 in func, param2 =", param2)
}
