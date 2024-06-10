package main

import (
	"context"
	"fmt"
	"time"
)

func interface1(ctx context.Context) error {
	// 模拟接口1的调用
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Interface 1 call completed successfully")
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func interface2(ctx context.Context) error {
	// 模拟接口2的调用
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Interface 2 call completed successfully")
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func main() {
	rpcTimeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	err := interface1(ctx)
	if err != nil {
		fmt.Println("Interface 1 call failed:", err)
	}

	// 使用同一个context，第一个rpc超时后，后续的rpc复用ctx直接超时
	start := time.Now()
	err = interface2(ctx)
	duration := time.Since(start)
	fmt.Printf("Execution time: %s\n", duration)
	if err != nil {
		fmt.Println("Interface 2 call failed:", err)
	}
}
