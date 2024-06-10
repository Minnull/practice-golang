package main

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// defer时修改退出的error状态
func main() {
	errFalse := printStr(false)
	fmt.Println(errFalse)

	errTrue := printStr(true)
	fmt.Println(errTrue)
}

func printStr(hasErr bool) (err error) {

	defer func() {
		err1 := status.Error(codes.ResourceExhausted, "")
		if err == nil {
			err = err1
		}
	}()

	fmt.Println("打印退出")
	if hasErr {
		return status.Error(codes.PermissionDenied, "")
	}

	return nil
}
