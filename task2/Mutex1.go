// 题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

// 考察点 ： sync.Mutex 的使用、并发数据安全。
package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// 添加计数器
func (s *SafeCounter) add() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++

}

// 获取计数器
func (s *SafeCounter) get() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}
func main() {
	SafeCounter := SafeCounter{}
	var wg sync.WaitGroup
	numGoroutines := 10
	incrementsPerGoroutine := 1000

	wg.Add(numGoroutines)

	fmt.Printf("启动 %d 个goroutine，每个进行 %d 次递增操作\n",
		numGoroutines, incrementsPerGoroutine)
	fmt.Printf("期望结果: %d\n", numGoroutines*incrementsPerGoroutine)

	// 启动goroutine
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < incrementsPerGoroutine; j++ {
				SafeCounter.add()
			}

			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}

	// 等待所有goroutine完成
	wg.Wait()

	// 输出最终结果
	finalCount := SafeCounter.get()
	fmt.Printf("\n最终计数: %d\n", finalCount)
}
