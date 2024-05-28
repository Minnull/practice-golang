package main

import (
	"fmt"
	"sync/atomic"
)

var (
	subTaskBatchNum = make(map[string]*int64)
)

func getOrCreateDefaultValue(key string, data map[string]*int64) *int64 {
	if dataValue, exists := data[key]; exists {
		// 指针获取值，然后比较大小
		if *dataValue < 0 {
			atomic.StoreInt64(dataValue, 0)
		}
		return dataValue
	}
	int64Temp := int64(0)
	data[key] = &int64Temp
	return &int64Temp
}

func main() {
	value := getOrCreateDefaultValue("test", subTaskBatchNum)
	fmt.Println(*value)
	// 在外部，修改同一个内存地址，值修改生效
	atomic.StoreInt64(value, 100)
	value3 := getOrCreateDefaultValue("test", subTaskBatchNum)
	fmt.Println(*value3)
	atomic.StoreInt64(value, -100)
	value2 := getOrCreateDefaultValue("test", subTaskBatchNum)
	fmt.Println(*value2)
}
