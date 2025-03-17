package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 创建一个带超时时间的 context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 在 6 秒后取消上下文
	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("Canceling context from main")
		cancel()
	}()

	process(ctx)
}

func process(ctx context.Context) {
	var ctx2 context.Context
	var cancelFunc context.CancelFunc

	// 检查 ctx 是否有截止时间
	deadline, ok := ctx.Deadline()
	if ok {
		// 如果有截止时间，使用相同的截止时间
		ctx2, cancelFunc = context.WithDeadline(ctx, deadline)
	} else {
		// 如果没有截止时间，设置一个自定义的超时时间
		ctx2, cancelFunc = context.WithTimeout(ctx, 3*time.Second)
	}
	defer cancelFunc()

	// 在此处使用 ctx2
aa:
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("Operation completed within time")
		case <-ctx2.Done():
			fmt.Println("Operation canceled or timed out:", ctx2.Err())
			break aa
		}
	}
}
