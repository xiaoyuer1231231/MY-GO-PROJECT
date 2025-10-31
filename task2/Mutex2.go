// 题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。

//     考察点 ：原子操作、并发数据安全。

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 定义一个结构体
type SafeCounterone struct {
	count int64
}

func (s *SafeCounterone) add() {

	atomic.AddInt64(&s.count, 1)
}

func (s *SafeCounterone) get() int64 {
	return atomic.LoadInt64(&s.count)
}

func main() {
	SafeCounter := SafeCounterone{}
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()

			for j := 0; j < 1000; j++ {
				SafeCounter.add()
			}

			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}
	wg.Wait()
	resultCount := SafeCounter.get()
	fmt.Println(resultCount)
}
