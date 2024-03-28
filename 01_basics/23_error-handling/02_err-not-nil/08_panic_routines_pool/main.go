package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"
)

func main() {
	antPool, _ := ants.NewPool(10)

	wg := sync.WaitGroup{}
	wg.Add(2)

	antPool.Submit(func() {
		countNum(&wg)
	})

	antPool.Submit(func() {
		printPanic(&wg)
	})

	wg.Wait()
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

func countNum(sync *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Println("输出：", i)
		time.Sleep(time.Second)
	}
	sync.Done()
}
