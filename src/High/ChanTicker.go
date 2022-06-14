package High

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ChanTicker https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.5.md
func ChanTicker() {
	ticker := time.NewTicker(1e9)
	defer ticker.Stop()

	select {
	case u := <-ticker.C:
		fmt.Println(u)
	default:
		fmt.Println(123)
	}

	//timeDiffer := time.Duration(1) * time.Second
	//crontabJob(timeDiffer)

	//timeLimit()

	afterJob()
}

//定时任务的概念
func crontabJob(timeStamp time.Duration) {
	start := time.Now()
	chRate := time.Tick(timeStamp)
	for i := 1; i <= 5; i++ {
		<-chRate
		end := time.Now()
		fmt.Printf("运行时常 %s \n", end.Sub(start))
		fmt.Printf("当前是第%d次\n", i)
	}
}

func timeLimit() {
	tick := time.Tick(1e8)
	boom := time.Tick(5e8)

	for {
		select {
		case <-tick:
			fmt.Println("tick.....")
		case <-boom:
			fmt.Println("boom.....")
		default:
			fmt.Println("灭火了！")
			time.Sleep(5e7)
		}
	}
}

func afterJob() {
	ch := make(chan interface{}, 1)

	//异步请求接口
	go func() {
		httpClient := &http.Client{}
		request, _ := http.NewRequest("GET", "https://dove-test.rp-field.com", nil)
		res, _ := httpClient.Do(request)
		body, _ := ioutil.ReadAll(res.Body)
		var data interface{}
		json.Unmarshal([]byte(body), &data)
		ch <- data
	}()

L:
	for {
		select {
		case res := <-ch:
			fmt.Println(res)
		case <-time.After(1e9):
			fmt.Println("停留1s并执行结束")
			break L
		}
	}
}
