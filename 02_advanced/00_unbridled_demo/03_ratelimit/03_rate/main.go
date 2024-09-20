package main

import (
	"context"
	"log"

	"golang.org/x/time/rate"
)

func main() {
	ratePerSecond := 1 * 1024 * 1024 // 1MB 每秒
	burst := 2 * 1024 * 1024         // 允许 2MB 的突发流量
	limiter := rate.NewLimiter(rate.Limit(ratePerSecond), burst)

	for i := 1; i <= 5; i++ {
		dataSize := i * 3 * 1024 * 1024 // 模拟传输 i*3 MB 的数据，超出 burst

		log.Printf("开始处理 %d MB 的数据\n", dataSize/1024/1024)

		for dataSize > 0 {
			// 每次传输的最大值不能超过 burst
			toSend := burst
			if dataSize < burst {
				toSend = dataSize
			}

			// 使用限流器等待处理该批次数据
			err := limiter.WaitN(context.Background(), toSend)
			if err != nil {
				log.Fatalf("限流器出错: %v", err)
			}

			log.Printf("处理了 %d MB 的数据\n", toSend/1024/1024)

			// 更新剩余数据量
			dataSize -= toSend
		}
	}

	log.Println("所有数据处理完毕")
}
