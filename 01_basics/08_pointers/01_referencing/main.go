package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a -", a)
	fmt.Println("aa -", &a)

	var b = &a

	fmt.Println("b - ", b)

	a = 46
	fmt.Println("bb - ", b)
}
