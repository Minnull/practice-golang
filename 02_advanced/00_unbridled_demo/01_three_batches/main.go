package main

import "fmt"

// 把正整数分3批，数组里返回每批的最后一个数字
func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("打印第 %d:", i)
		fmt.Println(splitIntoBatches(i))
	}
}

func splitIntoBatches(num int) []int {
	if num <= 0 {
		return []int{0, 0, 0}
	}

	third := num / 3 // 每个条件的数量
	firstEnd := third
	secondEnd := 2 * third

	if secondEnd > num {
		secondEnd = num
	}
	if firstEnd > secondEnd {
		firstEnd = secondEnd
	}

	return []int{firstEnd, secondEnd, num}
}
