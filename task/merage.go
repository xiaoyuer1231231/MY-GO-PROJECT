package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println("currentInterval", intervals)

	merged := [][]int{}
	currentInterval := intervals[0]

	for _, interval := range intervals[1:] {
		fmt.Println("interval", interval[0])
		fmt.Println("currentInterval", currentInterval[1])

		if interval[0] <= currentInterval[1] {
			// 有重叠，合并区间
			currentInterval[1] = max(currentInterval[1], interval[1])
			fmt.Println("maxcurrentInterval", currentInterval[1])

		} else {
			fmt.Println("elsecurrentInterval", currentInterval)

			// 无重叠，添加当前区间到结果中
			merged = append(merged, currentInterval)
			currentInterval = interval
		}
	}
	// 添加最后一个区间
	merged = append(merged, currentInterval)
	return merged
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	intervals := [][]int{{1, 3}, {3, 8}, {2, 6}, {15, 18}}
	result := merge(intervals)
	fmt.Println("Merged Intervals:", result)
}
