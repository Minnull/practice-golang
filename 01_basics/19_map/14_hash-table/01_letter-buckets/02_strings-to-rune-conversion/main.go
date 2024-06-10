package main

import "fmt"

func main() {
	// 将字符串 "A" 转换为 rune 类型，取其第一个字符，即 'A' 的 Unicode 编码值
	letter := rune("A"[0])
	// 打印输出 'A' 的 Unicode 编码值
	fmt.Println(letter)
}
