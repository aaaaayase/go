package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "rpc-first/hello-server/proto"
)

type server struct {
	pb.UnimplementedSayHelloServer
}

func (c *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: req.Name}, nil
}

func main() {
	// 开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	newServer := grpc.NewServer()
	// 在grpc服务中去注册我们自己的服务
	pb.RegisterSayHelloServer(newServer, &server{})
	// 启动服务
	err := newServer.Serve(listen)
	if err != nil {
		fmt.Println("failed to serve:%v", err)
		return
	}
}
