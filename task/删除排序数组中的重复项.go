package main

import "fmt"

func main() {
	ss := []int{1, 2, 2, 3, 4, 5, 5, 6, 6, 7, 8, 8}
	result := arr(ss)
	fmt.Print(result)

}
func arr(aa []int) []int {
	n := len(aa)
	if len(aa) == 0 {
		return aa
	}
	var i int = 1
	for k := 1; k < n; k++ {
		if aa[k] != aa[k-1] {

			aa[i] = aa[k]

			i++
		}
	}
	return aa[:i]
}
