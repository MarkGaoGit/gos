package main

import (
	"fmt"
	"gos/src/c4"
	"time"
)

func main() {
	s := time.Now()
	c4.Fibonacci()
	e := time.Now()

	c := e.Sub(s)
	fmt.Println(c)

	// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.0.md
}