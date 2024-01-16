package main

import (
	"fmt"
	"sync"
)

//为什么最终没有输出?

//最终没有输出的原因是因为主 goroutine 在执行 for n := range c 循环时，会一直等待从通道 c 中接收值，直到通道被关闭。然而，在这个代码中，通道 c 是在另一个 goroutine 中关闭的，而主 goroutine 在等待接收值时并不知道通道何时会被关闭。
//具体来说，当主 goroutine 执行到 for n := range c 时，它会尝试从通道 c 中接收值。如果通道 c 是未关闭的，并且没有发送方向通道发送值，那么主 goroutine 会阻塞在这里等待值的到来。但是在这个代码中，通道 c 是在另一个 goroutine 中关闭的，而主 goroutine 并不知道何时会有值被发送到通道中。
//因此，当通道 c 在另一个 goroutine 中关闭时，主 goroutine 仍然在等待接收值的操作中阻塞。这就导致了程序无法继续执行下去，也就没有输出。
//要解决这个问题，可以使用 sync.WaitGroup 来等待所有的 goroutine 完成任务后再关闭通道。在每个匿名函数中，通过 wg.Add(1) 增加等待组的计数器，然后在任务完成时调用 wg.Done() 来减少计数器。最后，在最后一个匿名函数中使用 wg.Wait() 来等待所有的任务完成后再关闭通道。这样可以确保在关闭通道之前，所有的发送操作都已经完成，避免主 goroutine 在等待接收值时被阻塞。

func main() {

	c := make(chan int)

	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i < 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
