package main

import (
	"context"
	"fmt"
	"time"
)

func interface1(ctx context.Context) error {
	// 模拟接口1的调用
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Interface 1 call completed successfully")
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func interface2(ctx context.Context) error {
	// 模拟接口2的调用
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Interface 2 call completed successfully")
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func main() {
	rpcTimeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	err := interface1(ctx)
	if err != nil {
		fmt.Println("Interface 1 call failed:", err)
	}

	// 创建新的上下文以供接口2使用
	ctx2, cancel2 := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel2()

	start := time.Now()
	err = interface2(ctx2)
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
	if err != nil {
		fmt.Println("Interface 2 call failed:", err)
	}
}
