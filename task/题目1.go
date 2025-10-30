// 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
package main

import "fmt"

func array(a [10]int) {
	s := make([]int, 0)
	seen := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		fmt.Println("use len(), index = ", i, "value = ", a[i])
		for j := i + 1; j < len(a); j++ {
			fmt.Println("use len(), index = ", j, "value = ", a[j])
			if a[i] == a[j] && !seen[a[i]] {
				// 保存数据
				s = append(s, a[i])
				seen[a[i]] = true
				break
			}

		}
	}
	fmt.Println("结果", s)
}
func findSingleNumber(a [10]int) []int {
	countMap := make(map[int]int)
	s := make([]int, 0)

	for i := 0; i < len(a); i++ {
		countMap[a[i]]++
	}
	fmt.Println("只出现一次的数字是:", countMap)

	for num, count := range countMap {
		if count == 1 {
			// 添加数据
			s = append(s, num)
		}
	}

	return s // 如果没有找到，返回-1
}
func main() {
	var a [10]int = [10]int{1, 2, 3, 4, 1, 2, 2, 1, 9, 6}
	// array(a)
	result := findSingleNumber(a)
	fmt.Println("只出现一次的数字是:", result) // 输出: 只出现一次的数字是: 3
}
