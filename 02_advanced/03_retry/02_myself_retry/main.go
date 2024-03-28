package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	result, err := customRetry(performTaskError, 3, 1*time.Second)
	if err != nil {
		fmt.Println("Function failed after multiple retries:", err)
	} else {
		fmt.Println("Function succeeded! Result:", result)
	}
}

func customRetry(fn interface{}, maxAttempts int, delay time.Duration) (interface{}, error) {
	for attempt := 1; attempt <= maxAttempts; attempt++ {

		f, ok := fn.(func() (interface{}, error))
		if !ok {
			return nil, errors.New("invalid function type")
		}

		result, err := f()
		if err == nil {
			return result, nil
		}

		fmt.Printf("Error occurred: %s. Retrying...\n", err.Error())
		time.Sleep(delay)
	}

	return nil, errors.New("function failed after multiple retries")
}

func performTaskError() (interface{}, error) {
	fmt.Println("执行外部结果")
	return nil, errors.New("task failed")
}
