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

func testRateLimiter(rate, dataSize int64) {
	limiter := NewRateLimiter(rate)

	log.Printf("开始处理 %d MB 的数据\n", dataSize/1024/1024)

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
