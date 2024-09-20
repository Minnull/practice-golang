package main

import (
	"log"
	"sync"
	"time"
)

// 限流器
type RateLimiter struct {
	mu       sync.Mutex
	rate     int64 // 每秒允许的字节数
	interval time.Duration
	ticker   *time.Ticker
	stop     chan struct{}
	bucket   int64
}

func NewRateLimiter(rate int64) *RateLimiter {
	interval := time.Second
	rl := &RateLimiter{
		rate:     rate,
		interval: interval,
		ticker:   time.NewTicker(interval),
		stop:     make(chan struct{}),
		bucket:   rate,
	}
	go rl.refill()
	return rl
}

func (rl *RateLimiter) refill() {
	for {
		select {
		case <-rl.ticker.C:
			rl.mu.Lock()
			rl.bucket = rl.rate
			rl.mu.Unlock()
		case <-rl.stop:
			return
		}
	}
}

func (rl *RateLimiter) Allow(size int64) {
	for {
		rl.mu.Lock()
		if rl.bucket >= size {
			rl.bucket -= size
			rl.mu.Unlock()
			return
		}
		rl.mu.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}

func (rl *RateLimiter) UpdateRate(newRate int64) {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.rate = newRate
	rl.bucket = newRate

	close(rl.stop)
	rl.ticker.Stop()

	rl.stop = make(chan struct{})
	rl.ticker = time.NewTicker(rl.interval)
	go rl.refill()
}

// 超过限流最大值，会一直卡死
func main() {
	limiter := NewRateLimiter(1 * 1024 * 1024) // 1MB/s

	// 模拟数据传输
	go func() {
		for i := 0; i < 1000; i++ {
			limiter.Allow(1024 * 1024) // 1MB
			log.Printf("Sent data,num= %d", i)
		}
	}()

	// 模拟动态调整限流速率
	go func() {
		time.Sleep(5 * time.Second)
		log.Println("Updating rate limit to 2MB/s")
		limiter.UpdateRate(2 * 1024 * 1024)

		time.Sleep(5 * time.Second)
		log.Println("Updating rate limit to 4MB/s")
		limiter.UpdateRate(4 * 1024 * 1024)
	}()

	// 防止主协程退出
	select {}
}
