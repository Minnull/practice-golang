package main

import "fmt"

func main() {
	xs := []int{1, 3, 5, 7, 9, 11}

	for index, value := range xs {
		fmt.Println(index, "-", value)
	}
}
