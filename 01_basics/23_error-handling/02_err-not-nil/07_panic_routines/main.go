package main

import (
	"fmt"
	"sync"
)

func main() {

	group := sync.WaitGroup{}
	group.Add(1)
	go printPanic(&group)

	group.Wait()
}

func printPanic(sync *sync.WaitGroup) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			sync.Done()
		}
	}()

	fmt.Println("开始")
	panic("print panic")
	fmt.Println("结束")
	sync.Done()
}
