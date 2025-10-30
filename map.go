package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2) // 告诉 WaitGroup：需要等待 2 个 goroutine

	// 启动第一个 goroutine
	go func() {
		defer wg.Done() // 完成后调用 Done()
		fmt.Println("Goroutine 1 开始")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 1 结束")
	}()

	// 启动第二个 goroutine
	go func() {
		defer wg.Done() // 完成后调用 Done()
		fmt.Println("Goroutine 2 开始")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 2 结束")
	}()

	fmt.Println("主线程等待中...")
	wg.Wait() // 阻塞，直到 2 个 Done() 被调用
	fmt.Println("所有 goroutine 完成，主线程继续")
}
