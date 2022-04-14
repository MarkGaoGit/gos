package Web

import (
	"flag"
	"fmt"
	"net"
)

const maxRead = 512

func TcpMain() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("Error: 启动服务需要传入地址 和 端口号")
	}

	listenAddr := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	listener := initServer(listenAddr)
	for true {
		request, err := listener.Accept()
		checkError(err, "请求获取失败！")

		go connectHandle(request)
	}
}

// connectHandle 处理请求
func connectHandle(request net.Conn) {
	connForm := request.RemoteAddr().String()
	println("远程的连接地址来自：" + connForm)

	sayHello(request)

	for true {
		var ibuf = make([]byte, maxRead+1)
		length, err := request.Read(ibuf[0:maxRead])
		switch err {
		case nil:
			handleMsg(length, ibuf)
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err := request.Close()
	println("客户端服务关闭!")
	checkError(err, "客户端服务关闭！")
}

// sayHello 问候
func sayHello(request net.Conn) {
	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
	wrote, err := request.Write(obuf)
	checkError(err, "Write: wrote "+string(rune(wrote))+" bytes.")
}

// handleMsg 消息处理
func handleMsg(length int, ibuf []byte) {
	if length > 0 {
		fmt.Println(string(ibuf))
	}
}

// initServer 初始化tcp服务
func initServer(hostServer string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostServer)
	checkError(err, "服务器端口初始化失败: '"+hostServer+"'")
	listen, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "服务器端口启动失败！")
	println("服务端口启动完成: " + listen.Addr().String())
	return listen
}

func checkError(error error, info string) {
	if error != nil {
		panic("Error：" + info + "【" + error.Error() + "】")
	}
}
