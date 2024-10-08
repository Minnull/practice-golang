package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/Minnull/practice-golang/02_advanced/02_grpc/02_stream/pro"
	"google.golang.org/grpc"
)

const PORT = ":50051"

type GreeterServer struct {
	pro.UnimplementedGreeterServer
}

// 服务端 单向流
func (s *GreeterServer) GetStream(req *pro.StreamReqData, res pro.Greeter_GetStreamServer) error {
	i := 0
	log.Println("外部开始执行", i)
	for {
		i++
		log.Println("开始发送", i)
		// sleep可以减速，观察现象
		// 不管加sleep还是不加sleep：测试重启client，server不会立即停止发送，而是继续发送一段时间后停止。client重启后，无法断点续传，
		//
		// server stream1发送的数据：1、2、3、4、【这时候client停止服务，server没有感知，继续发送后面的数据】、5、6、7、8、9、10
		// server stream2发送的数据：1、2、3、4、5、6、7、8、9、10
		// client接收的数据：1、2、3、4、【client停止服务，停止接收stream1流的数据，重启后，后面是stream2发送的数据】、1、2、3、4、5、6、7、8、9、10
		//
		// 总结：client关闭不会影响server的stream流停止发送数据，client新创建连接后，会新建stream流从头开始接收数据。
		// 并且server的stream1、stream2流同时存在，继续发送数据，client只与最新的stream2保持数据同步
		//time.Sleep(3 * time.Second)
		res.Send(&pro.StreamResData{Data: fmt.Sprintf("%v %v", time.Now().Unix(), i)})
		log.Println("发送完成", i)
		time.Sleep(1 * time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

// 客户端 单向流
func (this *GreeterServer) PutStream(cliStr pro.Greeter_PutStreamServer) error {

	for {
		if tem, err := cliStr.Recv(); err == nil {
			log.Println(tem)
		} else {
			log.Println("break, err :", err)
			break
		}
	}

	return nil
}

// 客户端服务端 双向流
func (this *GreeterServer) AllStream(allStr pro.Greeter_AllStreamServer) error {

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			data, _ := allStr.Recv()
			log.Println(data)
		}
		wg.Done()
	}()

	go func() {
		for {
			allStr.Send(&pro.StreamResData{Data: "ssss"})
			time.Sleep(time.Second)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func main() {
	//监听端口
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		return
	}
	//创建一个grpc 服务器
	s := grpc.NewServer()
	//注册事件
	pro.RegisterGreeterServer(s, &GreeterServer{})
	//注意这里这个pro是你定义proto文件生成后的那个go文件中引用的，而后面的这个函数是注册名称，是根据你自己定义的名称生成的，不同的文件该函数名称是不一样的，不过都是register这个单词开头的，这里不能原样照搬
	//处理链接
	s.Serve(lis)
}
