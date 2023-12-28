package main

import (
	"fmt"
	"github.com/Minnull/practice-golang/02_advanced/01_rpc/02_standard/svc"
	"log"
)

func main() {
	client, err := svc.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("fmt print reply: " + reply)
}
