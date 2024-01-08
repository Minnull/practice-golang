package main

import "fmt"

func main() {
	for i := 100; i <= 200; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)
}
