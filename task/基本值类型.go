package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 8, 9, 9}
	s := plusOne(ints)
	fmt.Println(s)
}
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	digits = make([]int, n+1)
	return digits
}
