package main

import "fmt"

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	for i := 0; i < 10; i++ {
		if i >= 5 {
			panic("num > 4")
		}
		fmt.Println(i)
	}

	fmt.Println("结束")
}
