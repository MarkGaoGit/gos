package High

import "fmt"

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

func RequestLimit() {
	service := make(chan *Request)
	go serverHandle(service)
}

func serverHandle(service chan *Request) {
	for {
		request := <-service
		go handle(request)
	}
}

func handle(r *Request) {
	sem <- 1 //放什么都可以
	fmt.Println(r)
	<-sem // 移除后 进行下一个请求处理
}
