package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// 1.并行任务中，有1个task返回了业务自定义错误时，其他task已经在循环ctx信号了：
//   当其中一个任务返回业务错误后，该任务会尽快退出，同时如果你使用了 errgroup.WithContext，上下文也会被取消。
//   不过需要注意的是，g.Wait() 本身不会立刻返回错误，而是会等待所有已启动的 goroutine 退出后再返回。
//   这意味着，即使某个任务返回错误后，其他任务也会继续运行（或因 context 被取消而退出），直到所有任务结束后，g.Wait() 才会返回第一个遇到的错误。
//
// 2. 并行任务中，有1个 task 返回了业务自定义错误时，存在1个 task 还在执行，其他 task 已经在循环 ctx 信号了：
//   当某个 task 返回业务自定义错误后，该 task 会尽快退出，并且因为使用了 errgroup.WithContext，上下文会被取消。
//   取消信号会传递给其他 task，促使它们在下一个检测 ctx.Done() 时退出。
//   然而，如果有个 task 处于一个密集计算或者长时间阻塞的操作中，还未检测到 ctx 被取消（比如还在执行下一次循环或睡眠中），
//   它将继续执行一段时间，直到在下一个 select 轮次中检测到 ctx.Done() 信号后退出。
//   最终，errgroup 的 g.Wait() 会返回第一个捕获到的错误（业务自定义错误），但它会等待所有 task 退出后再返回这个错误。
//
// 3.任务运行中 context 被取消：
//   如果外部调用了 cancel() 导致 context 被取消，任务应当通过定期检查 ctx.Done() 来及时退出，并返回 ctx.Err()。
//   在这种情况下，g.Wait() 最终会返回 ctx.Err()，不过它同样会等待所有 goroutine 退出后再返回结果。

// 切换开关：
// 设置为 true 时，模拟业务错误（任务3 返回错误）；
// 设置为 false 时，模拟外部取消（2秒后主动 cancel）。
const testBusinessError = true

func main() {
	// case 1
	//tasks := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 11}
	// case 2
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	ctx, cancel := context.WithCancel(context.Background())
	g, ctx := errgroup.WithContext(ctx)

	for _, task := range tasks {
		taskID := task // 捕获变量，避免闭包问题
		g.Go(func() error {
			return A(ctx, taskID)
		})
	}

	// 如果不测试业务错误，则启动外部取消逻辑
	if !testBusinessError {
		go func() {
			// case 3
			time.Sleep(2 * time.Second)
			fmt.Println("主动触发取消...")
			cancel()
		}()
	}

	// 等待所有任务退出，返回第一个错误（或 nil）
	if err := g.Wait(); err != nil {
		fmt.Printf("任务被取消或出错: %v\n", err)
	} else {
		fmt.Println("所有任务执行成功")
	}
}

func A(ctx context.Context, taskID int) error {
	fmt.Printf("任务 %d: 开始执行 A 方法\n", taskID)

	// 模拟业务错误：如果开关打开且 taskID==3，则返回错误
	if testBusinessError && taskID == 3 {
		// case 1
		// 延时一点，确保其他任务已启动
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("任务 %d: 模拟业务error\n", taskID)
		return fmt.Errorf("任务 %d: A 方法失败", taskID)
	}

	if testBusinessError && taskID == 7 {
		// case 2
		time.Sleep(30 * time.Second)
		fmt.Printf("任务 %d: 模拟等待结束\n", taskID)
		return fmt.Errorf("任务 %d: A 方法失败", taskID)
	}

	// 模拟长时间运行任务，通过 select 监听 ctx.Done() 实现取消检测
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Printf("任务 %d: 正在运行\n", taskID)
		case <-ctx.Done():
			fmt.Printf("任务 %d: 被取消\n", taskID)
			return ctx.Err()
		}
	}
}
