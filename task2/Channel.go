package main

// 题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。

//     考察点 ：通道的基本使用、协程间通信 ,带缓冲的

import (
	"fmt"
	"sync"
)

func receiveOnly(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // 完成后通知WaitGroup
	for v := range ch {
		fmt.Printf("接收到: %d\n", v)
	}
	fmt.Println("接收协程结束")
}

// 只发送channel的函数
func sendOnly(ch chan<- int, wg *sync.WaitGroup) {
	var once sync.Once
	defer wg.Done() // 完成后通知WaitGroup
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Printf("发送: %d\n", i)
	}
	// 确保通道只被关闭一次
	once.Do(func() {
		close(ch)
		fmt.Println("通道已关闭")
	})
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 3)
	go sendOnly(ch, &wg)
	go receiveOnly(ch, &wg)
	wg.Wait() // 等待所有goroutine完成
	fmt.Println("程序结束")
}
