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

func main() {
	rpcTimeout := 1 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()

	err := interface1(ctx)
	if err != nil {
		fmt.Println("Interface 1 call failed:", err)
	}
}
