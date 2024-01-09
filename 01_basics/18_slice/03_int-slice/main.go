package main

import "fmt"

func main() {
	// 通过内置方法make()创建切片，第一个参数表示切片类型，第二个参数表示切片中的元素长度，第三个参数表示切片可分配的空间大小
	mySlice := make([]int, 0, 3)

	fmt.Println("----------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("----------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice[i])
	}
}
