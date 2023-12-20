package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprintln("n > 10")
	} else {
		return fmt.Sprintln("n < 10")
	}
}
