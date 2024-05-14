package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// 启动一个 goroutine 处理任务
	go func() {
		for {
			select {
			case <-ctx.Done():
				// 当 ctx 的 done 信号触发时退出循环
				fmt.Println("Context done, exiting...")
				return
			default:
				// 执行阻塞的逻辑
				fmt.Println("Blocking operation...")
				// select 无法退出这里的阻塞
				time.Sleep(100 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Cancel...")
	cancel()
	// 主 goroutine 等待一段时间
	time.Sleep(50000000 * time.Second)
	fmt.Println("Main goroutine exits.")
}
