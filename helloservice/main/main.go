package main

import (
	"grpcdemo/helloservice"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// NewServer 创建一个 gRPC 服务器，它没有注册服务，也没有开始接受请求。
	grpcServer := grpc.NewServer()
	// 注册服务
	helloservice.RegisterHelloServiceServer(grpcServer, new(helloservice.HelloService))

	// 开启一个tcp监听
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server started...")
	// 在监听器 listen 上接受传入的连接，创建一个新的ServerTransport 和 service goroutine。 服务 goroutine读取 gRPC 请求，然后调用注册的处理程序来回复它们。
	log.Fatal(grpcServer.Serve(listen))
}
