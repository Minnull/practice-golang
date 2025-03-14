package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 假设有多个任务需要并发执行
	tasks := []int{1, 2, 3, 4, 5}

	for _, task := range tasks {
		wg.Add(1)
		go func(task int) {
			defer wg.Done()
			A(task)
		}(task)
	}

	wg.Wait()
	fmt.Println("All tasks completed")
}

// A 方法包含 C 和 D 步骤
func A(task int) {
	fmt.Printf("Starting task %d\n", task)

	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		C(task)
	}()
	go func() {
		defer wg.Done()
		D(task)
	}()

	wg.Wait()
	fmt.Printf("Completed task %d\n", task)
}

// C 步骤
func C(task int) {
	time.Sleep(1 * time.Second) // 模拟耗时操作
	fmt.Printf("Task %d: C step completed\n", task)
}

// D 步骤
func D(task int) {
	time.Sleep(1 * time.Second) // 模拟耗时操作
	fmt.Printf("Task %d: D step completed\n", task)
}
