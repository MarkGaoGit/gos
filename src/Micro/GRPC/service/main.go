package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gos/src/Micro/GRPC/service/proto/pr1"
	"net"
)

type Hello struct{}

// SayHello 需要实现 pr1.proto中生成的interface 方法
// 具体的参数、返回、可以看proto生成的文件中的 sayHello方法
func (h *Hello) SayHello(c context.Context, req *pr1.HelloReq) (*pr1.HelloRes, error) {

	fmt.Println(req)

	return &pr1.HelloRes{
		Message: "入参>>>>>>>>" + req.Name,
	}, nil
}

func main() {

	//定义远程调用的结构体和具体的方法
	grpcServer := grpc.NewServer()

	//注册GRPC服务 需要调用生成好的protobuf文件
	pr1.RegisterPr1Server(grpcServer, new(Hello))

	//监听已注册的端口
	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("服务监听失败！" + err.Error())
	}

	grpcServer.Serve(listen)

}
