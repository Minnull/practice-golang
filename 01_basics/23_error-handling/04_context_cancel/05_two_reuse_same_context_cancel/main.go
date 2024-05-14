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
	// 使用同一个context，并且已执行了cancel，后续使用直接cancel
	cancel()

	err := interface1(ctx)
	if err != nil {
		fmt.Println("Interface 1 call failed:", err)
	}

	start := time.Now()
	err = interface2(ctx)
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
	if err != nil {
		fmt.Println("Interface 2 call failed:", err)
	}
}
