package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带有超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

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
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 主 goroutine 等待一段时间
	time.Sleep(10 * time.Second)
	fmt.Println("Main goroutine exits.")
}
