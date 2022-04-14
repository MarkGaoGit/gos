package Web

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

// TcpServerClient
//TCP服务
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.1.md
func TcpServerClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	checkError(err)

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("输入信息：")

	clientName, _ := inputReader.ReadString('\n')

	trimmedClient := strings.Trim(clientName, "\r\n")

	for true {
		fmt.Println("输入Q 或 q退出程序！")

		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		if strings.ToUpper(trimmedInput) == "Q" {
			return
		}

		_, err := conn.Write([]byte(trimmedClient + "【says】" + trimmedInput + "\n"))
		checkError(err)
	}

}

func checkError(error error) {
	if error != nil {
		panic("Error：" + error.Error())
	}
}

func SocketOpenRead() {
	var (
		host   = "www.jd.cn"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]uint8, 4096)
		read   = true
		count  = 0
	)

	conn, err := net.Dial("tcp", remote)
	if err != nil {
		fmt.Println("链接错误" + err.Error())
		return
	}
	io.WriteString(conn, msg)

	for read {
		count, err = conn.Read(data)
		read = err == nil
		fmt.Printf(string(data[0:count]))
	}
	conn.Close()

}

// lianxi 15.1
