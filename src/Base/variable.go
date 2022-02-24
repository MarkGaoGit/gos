package Base

import (
	"fmt"
	"os"
	"runtime"
)

func FcVariable() {
	var a, b *int
	var c float32
	var d string
	var e bool

	var (
		f       int
		g       string
		userAge = 30
		HOME    = os.Getenv("HOME")
		GOROOT  = os.Getenv("GOROOT")
	)

	user := "mark"

	fmt.Print(a)
	fmt.Print("\n")
	fmt.Print(b)
	fmt.Print("\n")
	fmt.Print(c)
	fmt.Print("\n")
	fmt.Println(d)
	fmt.Print("\n")
	fmt.Print(e)
	fmt.Print("\n")
	fmt.Print(f)
	fmt.Print("\n")
	fmt.Println(g)
	fmt.Print("\n")

	fmt.Print(userAge)
	fmt.Print("\n")
	fmt.Println(HOME)
	fmt.Print("\n")
	fmt.Println(GOROOT)
	fmt.Print("\n")
	fmt.Println(user)

	var goos string = runtime.GOOS
	fmt.Printf("This opeartion systen is %s \n", goos)

	path := os.Getenv("PATH")
	fmt.Printf("this is path %s \n", path)

	//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/04.4.md
	//4.4.5 init 函数

}

var tmpStr = "c"

func Example1() {
	print(tmpStr)
}

func Example2() {
	tmpStr := "d"
	print(tmpStr)
}

var a string

func F1() {
	a = "A"
	print(a)
	F2()
}

func F2() {
	a = "C"
	print(a)
	F3()
}

func F3() {
	print(a)
}
