package main

import "fmt"

// 在这种方式下，切片变量被声明但未初始化，其值为nil。需要注意的是，这种方式定义的切片无法直接使用，需要通过make函数进行初始化后才能使用。
func main() {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
