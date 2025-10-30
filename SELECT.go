package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "来自ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "来自ch2"
	}()

	// select 会等待第一个可用的case
	select {
	case msg1 := <-ch1:
		fmt.Println("收到:", msg1) // 1秒后执行这个
	case msg2 := <-ch2:
		fmt.Println("收到:", msg2)
	}
}
