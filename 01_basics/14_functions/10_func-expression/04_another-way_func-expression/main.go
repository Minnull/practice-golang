package main

import "fmt"

//func makeGreeter() func() string：这是一个函数声明，它定义了一个名为 makeGreeter 的函数。这个函数没有参数，返回类型为 func() string，也就是一个不接受参数并返回字符串类型的函数。
//
//return func() string { return "Hello world!" }：在 makeGreeter 函数中，它返回了一个匿名函数。这个匿名函数没有参数，直接返回字符串 "Hello world!"。

func makeGreeter() func() string {
	return func() string {
		return "Hello world!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
