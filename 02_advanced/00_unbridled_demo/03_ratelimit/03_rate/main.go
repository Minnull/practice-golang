package main

import (
	"log"
	"time"

	"golang.org/x/time/rate"
)

func testRateLimiter(ratePerSecond, burst, dataSize int) {
	limiter := rate.NewLimiter(rate.Limit(ratePerSecond), burst)

	log.Printf("开始处理 %d MB 的数据\n", dataSize/1024/1024)

	for dataSize > 0 {
		// 每次传输的最大值不能超过 burst
		toSend := burst
		if dataSize < burst {
			toSend = dataSize
		}

		// 使用限流器检查是否允许处理该批次数据
		if limiter.AllowN(time.Now(), toSend) {
			log.Printf("处理了 %d MB 的数据\n", toSend/1024/1024)
			dataSize -= toSend
		} else {
			// 如果不允许，则等待并重试
			time.Sleep(time.Second)
		}
	}

	log.Println("所有数据处理完毕")
}

func main() {
	// 测试不同的限流情况
	log.Println("Testing with rate < sendSize")
	testRateLimiter(1*1024*1024, 2*1024*1024, 3*1024*1024)
	time.Sleep(2 * time.Second)

	log.Println("Testing with rate = sendSize")
	testRateLimiter(3*1024*1024, 3*1024*1024, 3*1024*1024)
	time.Sleep(2 * time.Second)

	log.Println("Testing with rate > sendSize")
	testRateLimiter(4*1024*1024, 3*1024*1024, 3*1024*1024)

	log.Println("Testing with rate = 0 (should block indefinitely)")
	testRateLimiter(0, 2*1024*1024, 3*1024*1024)
	time.Sleep(2 * time.Second)

	log.Println("所有测试完毕")
}
