package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, current := range greeting {
		fmt.Println(i, current)
	}

	for i := 0; i < len(greeting); i++ {
		fmt.Println(greeting[i])
	}
}
