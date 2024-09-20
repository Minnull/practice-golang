package main

import (
	"log"
	"time"
)

// 限流器
type RateLimiter struct {
	rate      int64         // 每秒允许的字节数
	interval  time.Duration // 时间间隔
	ticker    *time.Ticker  // 定时器
	bytesSent int64         // 已发送的字节数
}

// 创建新的限流器
func NewRateLimiter(rate int64) *RateLimiter {
	interval := time.Second
	return &RateLimiter{
		rate:     rate,
		interval: interval,
		ticker:   time.NewTicker(interval),
	}
}

// 检查是否允许发送指定大小的数据
func (rl *RateLimiter) Allow(size int64) {
	rl.bytesSent += size
	if rl.bytesSent > rl.rate {
		<-rl.ticker.C
		rl.bytesSent = size
	}
}

// 测试限流器
func testRateLimiter(rate, dataSize int64) {
	limiter := NewRateLimiter(rate)

	log.Printf("开始处理 %d MB 的数据，限流为 %d MB/s\n", dataSize/1024/1024, rate/1024/1024)

	for dataSize > 0 {
		toSend := rate
		if dataSize < rate {
			toSend = dataSize
		}

		limiter.Allow(toSend)
		log.Printf("处理了 %d MB 的数据\n", toSend/1024/1024)
		dataSize -= toSend
	}

	log.Println("所有数据处理完毕")
}

func main() {
	// 测试不同的限流情况
	log.Println("Testing with rate < sendSize")
	testRateLimiter(1*1024*1024, 3*1024*1024)
	time.Sleep(2 * time.Second)

	log.Println("Testing with rate = sendSize")
	testRateLimiter(3*1024*1024, 3*1024*1024)
	time.Sleep(2 * time.Second)

	log.Println("Testing with rate > sendSize")
	testRateLimiter(4*1024*1024, 3*1024*1024)

	log.Println("Testing with rate = 0 (should block indefinitely)")
	testRateLimiter(0, 3*1024*1024)
	time.Sleep(2 * time.Second)
	log.Println("所有测试完毕")
}
