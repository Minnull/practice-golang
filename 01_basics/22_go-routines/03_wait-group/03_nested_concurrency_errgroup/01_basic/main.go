package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// case 1: 模拟手动触发cancel，任务主动退出
// case 2: 模拟不触发cancel，任务正常执行完成
func main() {
	// 示例任务列表
	tasks := []int{1, 2, 3, 4, 5, 6, 7}

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	for _, task := range tasks {
		taskID := task // 避免闭包问题
		g.Go(func() error {
			return A(ctx, taskID)
		})
	}

	// 模拟某个条件触发取消操作 (例如，5秒后手动触发)
	go func() {
		// case 1: 模拟手动触发cancel，任务主动退出
		//time.Sleep(3 * time.Second)
		// case 2: 模拟不触发cancel，任务正常执行完成
		time.Sleep(30 * time.Second)
		fmt.Println("主动触发取消...")
		cancel() // 主动触发退出
	}()

	// 等待所有任务完成
	if err := g.Wait(); err != nil {
		fmt.Printf("任务被取消或出错: %v\n", err)
	} else {
		fmt.Println("所有任务执行成功")
	}
}

func A(ctx context.Context, taskID int) error {
	fmt.Printf("任务 %d: 开始执行 A 方法\n", taskID)

	// 使用子 errgroup 并行执行 C 和 D
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return stepC(ctx, taskID)
	})

	g.Go(func() error {
		return stepD(ctx, taskID)
	})

	// 等待步骤完成
	if err := g.Wait(); err != nil {
		fmt.Printf("任务 %d: A 方法执行失败，错误: %v\n", taskID, err)
		return err
	}

	fmt.Printf("任务 %d: A 方法执行成功\n", taskID)
	return nil
}

func stepC(ctx context.Context, taskID int) error {
	fmt.Printf("任务 %d: 步骤 C 开始\n", taskID)
	select {
	case <-time.After(2 * time.Second):
		fmt.Printf("任务 %d: 步骤 C 完成\n", taskID)
		return nil
	case <-ctx.Done():
		fmt.Printf("任务 %d: 步骤 C 被取消\n", taskID)
		return ctx.Err()
	}
}

func stepD(ctx context.Context, taskID int) error {
	fmt.Printf("任务 %d: 步骤 D 开始\n", taskID)
	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("任务 %d: 步骤 D 完成\n", taskID)
		return nil
	case <-ctx.Done():
		fmt.Printf("任务 %d: 步骤 D 被取消\n", taskID)
		return ctx.Err()
	}
}
