package Base

import (
	"errors"
	"fmt"
)

func ErrorHandle() {
	fmt.Println("====================START====================")
	customErrorFunc()
	//结束程序的运行并抛出错误

	//panic("tmp没有第4个元素")
	recoverTest()

	fmt.Println("====================END======================")

}

func customErrorFunc() {
	customError := errors.New("参数错误：Params不可为空")

	fmt.Println(customError)
}

func badCall() {
	panic("This is badCall")
}

func recoverTest() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	badCall()
	fmt.Println("end recover func ")

}
