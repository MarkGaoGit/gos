package main

import (
	"fmt"
	"net"
	"net/rpc"
)

//RPC服务

type User struct{}

type UserCreate struct {
	Name string
	Age  int
}

type UserResponse struct {
	Success bool
	Message string
}

// UserCreate RPC调用的返回必须是一个指针类型
func (u *User) UserCreate(uc UserCreate, ucr *UserResponse) error {
	fmt.Println(uc)

	//gorm添加用户即可

	*ucr = UserResponse{
		Success: true,
		Message: "用户创建完成",
	}

	return nil
}

// http的方式src/Web/RpcServer.go

func main() {

	// 注册userService 微服务

	err := rpc.RegisterName("userService", new(User))
	if err != nil {
		fmt.Println("rpc服务注册失败" + err.Error())
		return
	}

	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("服务监听失败！" + err.Error())
	}

	defer listen.Close()

	//死循环 相当于不结束、 客户端请求后不会结束程序
	for true {
		fmt.Println("准备建立连接..... ")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("连接建立失败！" + err.Error())
		}

		go rpc.ServeConn(conn)
	}

}
