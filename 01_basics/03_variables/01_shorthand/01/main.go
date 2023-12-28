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

	fmt.Printf("%v \n", v1)
	fmt.Printf("%v \n", v2)
	fmt.Printf("%v \n", v3)
	fmt.Printf("%v \n", v4)
	fmt.Printf("%v \n", v5)
	fmt.Printf("%v \n", v6)
	fmt.Printf("%v \n", v7)
	fmt.Printf("%v \n", v8)
	fmt.Printf("%v \n", v9)

	// 进行匿名赋值，不指定变量类型，根据变量值的类型自动推断出变量类型
	var v10 = 1
	v11 := "str"
	fmt.Printf("%v \n", v10)
	fmt.Printf("%v \n", v11)
}
