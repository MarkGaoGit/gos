package High

import "fmt"

type Request struct {
	a, b   int
	replyc chan int
}

type binOp func(a, b int) int

// ServerAndClient 服务器和客户端之间的应用
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.10.md
func ServerAndClient() {
	adder, quit := startServer(func(a, b int) int { return a + b })
	const N = 10
	var reqs [N]Request

	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}

	for i := N - 1; i >= 0; i-- {
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail")
		} else {
			fmt.Printf("Request %d Success! \n", i)
		}
	}
	quit <- true

	fmt.Println("done!")
}

func startServer(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)

	go server(op, service, quit)
	return service, quit
}

func server(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}
