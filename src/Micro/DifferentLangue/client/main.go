package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type UserCreate struct {
	Name string
	Age  int
}

type UserResponse struct {
	Success bool
	Message string
}

// JSON RPC客户端调用

// 把默认的rpc改成 jsonrpc
// rpc.Dial 改成了net.Dial
//增加基于json编码解码的rpc服务 36line
// client.call 改成了 新建立的client
// main
func main() {
	conn, err := net.Dial("tcp", "10.2.4.216:2000")
	if err != nil {
		fmt.Println("rpc服务连接失败！", err.Error())
	}

	defer conn.Close()

	//建立基于json rpc 解码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var response UserResponse

	//userService.UserCreate 调用了用户为服务下面的用户创建方法
	//RPC调用的返回必须是一个指针类型
	// RPC数据传输使用的是json
	// 使用nc -l 本地ip 端口 监听到发送的数据是 {"method":"userService.UserCreate","params":[{"Name":"mark","Age":33}],"id":0}
	err = client.Call("userService.UserCreate", UserCreate{
		Name: "mark",
		Age:  33,
	}, &response)
	if err == nil {
		fmt.Println(response)
	}

}
