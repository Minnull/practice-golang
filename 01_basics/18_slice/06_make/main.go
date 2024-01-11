package main

import "fmt"

// 长度（len）：切片的长度表示当前切片中实际包含的元素个数。切片的长度可以动态变化，当元素被追加到切片中时，长度会增加；当元素从切片中删除时，长度会减少。
// 容量（cap）：切片的容量是指底层数组中切片可以容纳的最大元素个数。容量取决于切片的起始位置和底层数组的长度。
// 当切片的长度等于容量时，再追加元素会导致切片扩容，底层数组会重新分配更大的内存空间。
func main() {
	// 3 is length & capacity
	customerNumber := make([]int, 3)
	fmt.Println(len(customerNumber))
	fmt.Println(cap(customerNumber))

	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	greeting := make([]string, 3, 5)
	fmt.Println("原始容量")
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"
	//greeting[3] = "two" // error
	greeting = append(greeting, "dddd")

	fmt.Println(greeting[2])
	fmt.Println(greeting[3])
	fmt.Println("扩容后容量")
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

	greeting = append(greeting, "dddd", "test", "first")
	fmt.Println("连续扩容后容量")
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
