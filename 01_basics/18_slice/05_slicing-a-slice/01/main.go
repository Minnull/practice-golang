package main

import "fmt"

func main() {

	var results []int
	fmt.Println(results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	// 索引从0开始，打印索引2到3，包括2，不包括4
	fmt.Println(mySlice[2:4])
	// 打印索引2
	fmt.Println(mySlice[2])
	// 输出myString的索引为2的字节数
	fmt.Println("myString"[2])
	// 输出myString的索引为2的字符
	fmt.Println(string("myString"[2]))
}
