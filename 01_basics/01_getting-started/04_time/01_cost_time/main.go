package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// 这里放置需要测量耗时的代码
	// 模拟一段耗时的操作
	time.Sleep(2738 * time.Millisecond)

	durationInMilliseconds := time.Since(start).Milliseconds()

	// 将耗时以毫秒为单位保存为 int64 类型
	durationMillisecondsInt := int64(durationInMilliseconds)

	fmt.Printf("代码执行耗时（毫秒）：%d\n", durationMillisecondsInt)
}
