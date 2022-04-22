package Web

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	N, M int
}

func (a *Args) Multiply(reply *int) error {
	*reply = a.N * a.M
	return nil
}

func RPCServer() {

	cal := new(Args)
	rpc.Register(cal)
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		fmt.Println("服务启动失败")
	}

	go http.Serve(listener, nil)

	time.Sleep(1000e9)
}
