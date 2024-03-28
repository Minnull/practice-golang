package main

import "fmt"

func main() {
	switch "hello" {
	case "a":
		fmt.Println("a")
	case "b":
		fmt.Println("b")
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("default hello")
	}
}
