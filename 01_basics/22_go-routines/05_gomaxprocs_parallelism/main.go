package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 在 Go 语言中，每个包可以包含一个名为 `init` 的特殊函数。`init` 函数在程序运行时自动被调用，而不需要显式地调用它。
// 在给定的代码示例中，`init` 函数被定义在 `main` 包中，并且没有参数和返回值。因此，它符合了 `init` 函数的定义。
// 在 `main` 包中，当程序启动时，Go 运行时会自动调用 `init` 函数。这意味着在 `main` 函数执行之前，`init` 函数会被自动执行。
// 在这个特定的代码示例中，`init` 函数被用来设置并发执行的最大 CPU 核心数，通过 `runtime.GOMAXPROCS(runtime.NumCPU())` 这一行代码来实现。
// `runtime.NumCPU()` 函数返回当前计算机的 CPU 核心数，而 `runtime.GOMAXPROCS()` 函数设置并发执行的最大 CPU 核心数。
// 因此，在这个代码示例中，`init` 函数在程序运行时会在 `main` 函数执行之前被自动调用，用于设置并发执行的最大 CPU 核心数。
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
