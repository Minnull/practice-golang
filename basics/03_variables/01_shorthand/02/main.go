package main

import "fmt"

func main() {

	// 变量声明
	var v1 int
	var v2 float32
	var v3 [10]int
	var v4 []float32
	var v5 struct {
		age int
	}
	var v6 *int
	// 声明一个字典变量，默认值nil
	var v7 map[string]string
	// 声明一个方法变量，默认值nil
	var v8 func(x int) int
	// 声明一个接口变量，默认值nil
	var v9 interface{}

	fmt.Printf("%T \n", v1)
	fmt.Printf("%T \n", v2)
	fmt.Printf("%T \n", v3)
	fmt.Printf("%T \n", v4)
	fmt.Printf("%T \n", v5)
	fmt.Printf("%T \n", v6)
	fmt.Printf("%T \n", v7)
	fmt.Printf("%T \n", v8)
	fmt.Printf("%T \n", v9)

	// 进行匿名赋值，不指定变量类型，根据变量值的类型自动推断出变量类型
	var v10 = 1
	v11 := "str"
	fmt.Printf("%T \n", v10)
	fmt.Printf("%T \n", v11)
}
