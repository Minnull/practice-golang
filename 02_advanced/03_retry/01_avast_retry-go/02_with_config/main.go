package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/avast/retry-go"
)

func main() {
	err := retry.Do(
		func() error {
			err := performTask()
			if err != nil {
				fmt.Printf("Error occurred: %s. Retrying...\n", err.Error())
			}
			return err
		},
		retry.Attempts(3),
		retry.Delay(1*time.Second),
		retry.MaxJitter(100*time.Millisecond),
		retry.OnRetry(func(n uint, err error) { fmt.Printf("Retrying... (attempt %d)\n", n) }),
		retry.RetryIf(func(err error) bool { return err != nil }),
	)

	if err != nil {
		fmt.Println("Function failed after multiple retries:", err)
	} else {
		fmt.Println("Function succeeded!")
	}
}

func performTask() error {
	// 这里是任务逻辑的实现
	// 如果任务成功完成，返回 nil；如果发生错误，返回相应的错误
	return errors.New("task failed")
}
