package main

import (
	"context"
	"fmt"
	"grpcdemo/helloservice"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	// 连接grpc服务端
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 一元rpc
	unaryRpc(conn)
	// 流式rpc
	streamRpc(conn)

}

func unaryRpc(conn *grpc.ClientConn) {
	// 创建grpc客户端
	client := helloservice.NewHelloServiceClient(conn)
	// 发送请求
	reply, err := client.Hello(context.Background(), &helloservice.String{Value: "hello 666666~"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("unaryRpc recv: ", reply.Value)
}

func streamRpc(conn *grpc.ClientConn) {
	// 创建grpc客户端
	client := helloservice.NewHelloServiceClient(conn)
	// 生成ClientStream
	stream, err := client.Channel(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			// 发送消息
			if err := stream.Send(&helloservice.String{Value: "hi,I am Shuang!"}); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		// 接收消息
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println("streamRpc recv: ", recv.Value)

	}
}
