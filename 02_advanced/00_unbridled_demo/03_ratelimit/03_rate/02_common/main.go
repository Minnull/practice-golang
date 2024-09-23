package main

import (
	"context"
	"errors"
	"log"
	"time"

	"golang.org/x/time/rate"
)

// RateLimiter 包装
type RateLimiter struct {
	limiter *rate.Limiter
}

// NewRateLimiter 初始化 RateLimiter 对象
func NewRateLimiter(ratePerSecond, burst int) *RateLimiter {
	return &RateLimiter{
		limiter: rate.NewLimiter(rate.Limit(ratePerSecond), burst),
	}
}

// Allow 封装的方法，检查是否允许处理指定大小的数据
func (rl *RateLimiter) Allow(ctx context.Context, size int) bool {
	if err := rl.limiter.WaitN(ctx, size); err != nil {
		if errors.Is(err, context.Canceled) {
			log.Printf("上下文取消: %v\n", err)
		} else if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("上下文超时: %v\n", err)
		} else {
			log.Printf("等待时发生错误: %v\n", err)
		}
		return false
	}
	return true
}

// ProcessData 使用限流器处理数据
func (rl *RateLimiter) ProcessData(ctx context.Context, dataSize, burst int) {
	log.Printf("开始处理 %d MB 的数据\n", dataSize/1024/1024)

	for dataSize > 0 {
		// 每次传输的最大值不能超过 burst
		toSend := burst
		if dataSize < burst {
			toSend = dataSize
		}

		// 使用限流器检查是否允许处理该批次数据
		if rl.Allow(ctx, toSend) {
			log.Printf("处理了 %d MB 的数据\n", toSend/1024/1024)
			dataSize -= toSend
		} else {
			// 如果不允许，则等待并重试
			time.Sleep(time.Second)
		}
	}

	log.Println("所有数据处理完毕")
}

// testRateLimiter 测试限流器
func testRateLimiter(ratePerSecond, burst, dataSize int) {
	ctx := context.Background()
	limiter := NewRateLimiter(ratePerSecond, burst)
	limiter.ProcessData(ctx, dataSize, burst)
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
	time.Sleep(2 * time.Second)

	log.Println("Testing with rate = 0 (should block indefinitely)")
	testRateLimiter(0, 2*1024*1024, 3*1024*1024)
	time.Sleep(2 * time.Second)

	log.Println("所有测试完毕")
}
