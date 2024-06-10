package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	ch1 := make(chan int)

	// 启动一个 goroutine 处理任务
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context done")
			case ch1 <- 1:
				// 从 ch1 接收数据
				fmt.Println("Received from ch1")
			case ch1 <- 2:
				// 从 ch2 接收数据
				fmt.Println("Received from ch2")
			}
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Cancel")
	cancel()
	time.Sleep(5 * time.Second)
	fmt.Println("Main goroutine exits.")
}
