package main

import (
	"fmt"
)

func main() {
	// 创建一个通道
	ch := make(chan int, 2)

	// 将通道同时赋值给两个变量
	ch1 := ch
	ch2 := ch

	// 打印通道的地址
	fmt.Printf("ch1 address: %p\n", ch1)
	fmt.Printf("ch2 address: %p\n", ch2)

	ch <- 1
	fmt.Printf("ch1 len: %p\n", len(ch1))
	fmt.Printf("ch2 len: %p\n", len(ch2))
}
