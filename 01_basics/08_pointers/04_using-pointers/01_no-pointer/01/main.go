package main

import "fmt"

func zero(z int) {
	z = 0
}

func main() {
	x := 5
	zero(5)
	fmt.Println(x) // 5
}
