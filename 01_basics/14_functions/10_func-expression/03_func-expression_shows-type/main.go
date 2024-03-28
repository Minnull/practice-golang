package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
	fmt.Println("%T\n", greeting)
}
