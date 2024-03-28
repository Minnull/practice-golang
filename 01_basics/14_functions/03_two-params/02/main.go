package main

import "fmt"

func main() {
	greet("jane", "doe")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
