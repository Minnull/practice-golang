package main

import (
	"context"
	"fmt"
	"time"
)

func interface1(ctx context.Context) error {
	// 模拟接口1的调用
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Interface 1 call completed successfully")
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func interface2(ctx context.Context) error {
	// 模拟接口2的调用
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Interface 2 call completed successfully")
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func main() {
	rpcTimeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	// 使用同一个context，第一个rpc未超时，则后续的rpc可以复用ctx
	// 但是两次查询，使用同一个context超时时间累计。
	// interface1耗时 + interface2耗时 > rpcTimeout则超时
	start1 := time.Now()
	err := interface1(ctx)
	if err != nil {
		fmt.Println("Interface 1 call failed:", err)
	}
	duration1 := time.Since(start1)
	fmt.Printf("Execution time1: %s\n", duration1)

	start := time.Now()
	err = interface2(ctx)
	duration := time.Since(start)
	fmt.Printf("Execution time2: %s\n", duration)
	if err != nil {
		fmt.Println("Interface 2 call failed:", err)
	}
}
