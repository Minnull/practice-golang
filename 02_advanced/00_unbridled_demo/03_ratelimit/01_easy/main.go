package main

import (
	"log"
	"math"
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

func main() {
	limiter := NewRateLimiter(1 * 1024 * 1024) // 1MB/s

	// 模拟数据传输
	go func() {
		for i := 0; i < math.MaxInt64; i++ {
			limiter.Allow(300 * 1024) // 1MB
			log.Printf("Sent data,num= %d", i)
		}
	}()

	// 防止主协程退出
	select {}
}
