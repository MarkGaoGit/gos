package High

import (
	"flag"
	"fmt"
)

var in = flag.Int("n", 10, "输入线程数量")

func f(left, right chan int) {
	left <- 1 + <-right
}

// LianSProcess
// 链式协程
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.12.md
func LianSProcess() {
	flag.Parse()
	leftMost := make(chan int)
	var left, right chan int = nil, leftMost
	for i := 0; i < *in; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}

	right <- 0
	x := <-leftMost
	fmt.Println(x)
}
