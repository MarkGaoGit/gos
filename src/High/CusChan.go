package High

import (
	"fmt"
	"strconv"
	"time"
)

func MeChan() {
	//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md 14.2.7
	ch := make(chan int, 2)
	defer close(ch)

	go setData(ch)
	go getData(ch)

	arr := []int{1, 2, 3, 4, 5}
	go sum(arr, ch)
	sum := <-ch
	fmt.Printf("求和结果%d\n", sum)

	//工厂模式
	stream := chanFactory()
	go popFactory(stream)

	sendChan := make(chan int)
	receiveChan := make(chan string)

	//输入 输出
	//go setData(sendChan)
	go processChannel(sendChan, receiveChan)
	//go getData(receiveChan)

	//打印素数
	/*ch1 := make(chan int)
	go generate(ch1)
	for {
		prime := <-ch1
		fmt.Printf("输出的数 %d \n", prime)
		ch2 := make(chan int)
		go filter(ch1, ch2, prime)
		ch1 = ch2
	}*/

	time.Sleep(1e9)
}

func setData(ch chan int) {
	for i := 10; i <= 100; i += 10 {
		ch <- i
	}
}

func getData(ch chan int) {

	// 下面2中方法都可以从通道中获取到数据 也顺带了检测通道是否关闭
	/*for {
		num := <- ch
		fmt.Println(num)
	}*/

	for v := range ch {
		fmt.Println(v)
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func processChannel(in <-chan int, out chan<- string) {
	for inValue := range in {
		result := strconv.Itoa(inValue * 3)
		out <- result
	}
}

func chanFactory() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	return ch
}

func popFactory(ch chan int) {
	for {
		fmt.Printf("工厂通道输出 %d \n", <-ch)
	}
}

func sum(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}
