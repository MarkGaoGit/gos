package High

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	name  string
	value int
}

type Pool struct {
	Mu    sync.Mutex
	Tasks []*Task
}

// EventJob
//任务的分发
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.7.md
func EventJob() {
	pools := new(Pool)
	pools.Tasks = []*Task{
		{"job1", 1},
		{"job2", 2},
		{"job3", 3},
		{"job4", 4},
		{"job5", 5},
	}

	// 加锁的方式
	lockMethodHandle(pools)

	//协程通道处理方式
	sendCh, outCh := make(chan *Task), make(chan *Task)
	go func() {
		for _, v := range pools.Tasks {
			sendCh <- v
		}
	}()

	for i := 0; i < len(pools.Tasks); i++ {
		go chanMethodHandle(sendCh, outCh)
	}

	//要等协程运行结束
	time.Sleep(1e9)
}

func chanMethodHandle(sendCh, outCh chan *Task) {
	for {
		task := <-sendCh
		handleProcess(task)
		outCh <- task
	}
}

func lockMethodHandle(pools *Pool) {
	for {
		if len(pools.Tasks) < 1 {
			fmt.Println("任务分发完成！可休息 或 退出")
			break
		}
		pools.Mu.Lock()
		task := pools.Tasks[0]
		pools.Tasks = pools.Tasks[1:]
		pools.Mu.Unlock()
		handleProcess(task)
	}
}

func handleProcess(taskP *Task) {
	fmt.Printf("任务名称为%s, 值为%d \n", taskP.name, taskP.value)
}
