package Web

import (
	"fmt"
	"net/rpc"
)

// RPCClient https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.09.md
func RPCClient() {

	client, e := rpc.DialHTTP("tcp", "localhost:1234")
	if e != nil {
		fmt.Println("error service")
	}

	args := &Args{8, 7}

	var reply int
	e = client.Call("Args.Multiply", args, reply)
	fmt.Println("请求结果", reply)
}
