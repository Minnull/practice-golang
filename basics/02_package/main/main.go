package main

import (
	"fmt"
	"github.com/Minnull/practice-golang/basics/02_package/export_val"
	"github.com/Minnull/practice-golang/basics/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("123456"))
	fmt.Println(stringutil.MyName)
	fmt.Println(otherpackage.BearName)
}
