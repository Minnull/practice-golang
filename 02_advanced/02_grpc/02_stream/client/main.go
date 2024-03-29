package main

import (
	"context"
	"log"
	"time"

	"github.com/Minnull/practice-golang/02_advanced/02_grpc/02_stream/pro"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/balancer/grpclb"
)

// 参考代码：https://www.jianshu.com/p/bd35cbf279fb

const (
	ADDRESS = "localhost:50051"
)

func main() {
	//通过grpc 库 建立一个连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	//通过刚刚的连接 生成一个client对象。
	c := pro.NewGreeterClient(conn) //这个pro与服务端的同理，也是来源于proto编译生成的那个go文件内部调用
	//调用服务端推送流
	reqstreamData := &pro.StreamReqData{Data: "aaa"}
	res, _ := c.GetStream(context.Background(), reqstreamData)
	for {
		aa, err := res.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(aa)
	}
	//客户端 推送 流
	putRes, _ := c.PutStream(context.Background())
	i := 1
	for {
		i++
		putRes.Send(&pro.StreamReqData{Data: "ss"})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	//服务端 客户端 双向流
	allStr, _ := c.AllStream(context.Background())
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
	}()

	go func() {
		for {
			allStr.Send(&pro.StreamReqData{Data: "ssss"})
			time.Sleep(time.Second)
		}
	}()
	select {}
}
