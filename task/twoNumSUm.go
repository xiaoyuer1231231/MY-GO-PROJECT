package main

import "fmt"

func twoSumDebug(nums []int, target int) map[int]int {
	numMap := make(map[int]int)
	resultMap := make(map[int]int)

	fmt.Printf("开始两数之和搜索:\n")
	fmt.Printf("数组: %v, 目标: %d\n\n", nums, target)

	for i, num := range nums {
		complement := target - num
		fmt.Printf("步骤%d: 当前数字 nums[%d]=%d\n", i+1, i, num)
		fmt.Printf("      需要找: %d\n", complement)

		if index, found := numMap[complement]; found {
			fmt.Printf("      ✓ 找到匹配! nums[%d]=%d\n", index, complement)

			// 构建结果map
			resultMap[index] = complement
			resultMap[i] = num

			fmt.Printf("      构建结果map:\n")
			fmt.Printf("        resultMap[%d] = %d\n", index, complement)
			fmt.Printf("        resultMap[%d] = %d\n", i, num)
			fmt.Printf("      最终结果: %v\n", resultMap)

			return resultMap
		}

		numMap[num] = i
		fmt.Printf("      → 存入查找map: %d -> 索引%d\n", num, i)
		fmt.Printf("      当前查找map: %v\n\n", numMap)
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSumDebug(nums, 9)

	fmt.Printf("\n最终返回: %v\n", result)
}
