package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/avast/retry-go"
)

func main() {

	url := "http://baidu232323232.com"
	var body []byte

	err := retry.Do(
		func() error {
			fmt.Println("执行了")

			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return nil
		},
	)
	if err != nil {
		fmt.Println(" handle error")
	}

	fmt.Println(string(body))
}
