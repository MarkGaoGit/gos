package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gos/src/Micro/GRPC/client/proto/pr1"
)

func main() {

	//credentials.NewClientTLSFromFile() 从输入的证书文件中为客户端构造TLS凭证
	// grpc.WithTransportCredentials 配置连接级别的安全凭证 TLS/SSL等
	//建立tcp连接
	grpcClient, err := grpc.Dial("127.0.0.1:2000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("建立TCP连接失败" + err.Error())
	}

	//注册客户端

	client := pr1.NewPr1Client(grpcClient)
	res, err := client.SayHello(context.Background(), &pr1.HelloReq{
		Name: "二黑",
	})

	fmt.Println(res.GetMessage())
}
