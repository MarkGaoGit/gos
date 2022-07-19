package main

import (
	"fmt"
	"net/rpc"
)

type UserCreate struct {
	Name string
	Age  int
}

type UserResponse struct {
	Success bool
	Message string
}

// RPC客户端调用

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("rpc服务连接失败！")
	}

	defer client.Close()

	var response UserResponse

	//userService.UserCreate 调用了用户为服务下面的用户创建方法
	//RPC调用的返回必须是一个指针类型
	// RPC传输的数据使用的gob方式
	err = client.Call("userService.UserCreate", UserCreate{
		Name: "mark",
		Age:  33,
	}, &response)
	if err == nil {
		fmt.Println(response)
	}

}
