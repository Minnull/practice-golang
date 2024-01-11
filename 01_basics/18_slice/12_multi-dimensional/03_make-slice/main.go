package main

import "fmt"

// 在这种方式下，使用make函数创建了切片，并指定了切片的长度为35。切片被分配了底层数组，并且切片的长度和容量都被设置为35。这样就可以直接使用这些切片变量了。
func main() {
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
