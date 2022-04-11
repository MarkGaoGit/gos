package High

import (
	"fmt"
	"time"
)

// UseSelect https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.4.md
//select 做的就是：
//选择处理列出的多个通信情况中的一个。
//如果都阻塞了，会等待直到其中一个可以处理
//如果多个可以处理，随机选择一个
//如果没有通道操作可以处理并且写了 default 语句，它就会执行：default 永远是可运行的（这就是准备好了，可以执行）。
func UseSelect() {

	//使用select
	//useSelectFunc()

	//协程斐波那契数列
	processFibonacci()

	//协程生成随机数列 0 和 1
	processRandom01()

}

func processRandom01() {
	c := make(chan int)
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

}

func processFibonacci() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func useSelectFunc() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pumpF1(ch1)
	go pumpF2(ch2)
	go stuckF(ch1, ch2)

	time.Sleep(1e9)
}

func pumpF1(ch chan int) {
	for i := 1; ; i++ {
		ch <- i * 5
	}
}

func pumpF2(ch chan int) {
	for i := 2; ; i++ {
		ch <- i * 3
	}
}

func stuckF(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("通道【1】中出来的值: %d \n", v)
		case v := <-ch2:
			fmt.Printf("通道【2】中出来的值: %d \n", v)

		}
	}
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
