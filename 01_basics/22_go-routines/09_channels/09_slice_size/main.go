package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan []int, 1)

	group := sync.WaitGroup{}
	group.Add(2)

	go putData(ch, &group)
	go getData(ch, &group)

	group.Wait()
}

func putData(ch chan []int, group *sync.WaitGroup) {
	for i := 0; i < 100; i++ {

		slice := []int{1, 2, 3, 4, 5}
		slice = slice[:0]
		for j := 0; j <= i; j++ {
			slice = append(slice, j)
		}

		ch <- slice
		time.Sleep(1 * time.Second)
	}
	group.Done()
}

func getData(ch chan []int, group *sync.WaitGroup) {
	for i := 0; i < 100; i++ {
		fmt.Println(time.Now(), ":read data ", i, <-ch)
	}
	group.Done()
}
