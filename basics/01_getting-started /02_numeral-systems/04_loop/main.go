package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
