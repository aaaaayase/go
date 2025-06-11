package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "rpc-first/hello-server/proto"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("failed to dial:%v", err)
	}

	defer conn.Close()

	// 跟服务端建立连接
	client := pb.NewSayHelloClient(conn)
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Hello_world"})
	if err != nil {
		fmt.Println("failed to call:%v", err)
	}
	fmt.Println(resp.GetMessage())
}
