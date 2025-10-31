// 题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

// 考察点 ：协程原理、并发任务调度。
package main

import (
	"fmt"
	"sync"
	"time"
)

// 类型函数
type Task func()

func allTasks(tasks []Task) {
	var wg sync.WaitGroup
	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()
			start := time.Now()
			task()
			fmt.Printf("Task completed in %v\n", time.Since(start))
		}(task)
	}
	wg.Wait()
}

func main() {
	tasks := []Task{
		func() {
			for i := 0; i < 1; i++ {
				fmt.Println("Task 1 - Count:", i)
				time.Sleep(100 * time.Millisecond)
			}
		},
		func() {
			for i := 0; i < 1; i++ {
				fmt.Println("Task 2 - Count:", i)
				time.Sleep(200 * time.Millisecond)
			}
		},
		func() {
			for i := 0; i < 1; i++ {
				fmt.Println("Task 3 - Count:", i)
				time.Sleep(300 * time.Millisecond)
			}
		},
	}

	allTasks(tasks)
}
