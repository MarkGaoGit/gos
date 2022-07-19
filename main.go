package main

import (
	"fmt"
	"gos/src/Base"
	"time"
)

func init() {
	fmt.Println("==============program start!==============")
}

// https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/16.0.md
func main() {
	start := time.Now()

	//Web.ShiyongFunc()

	Base.ReadCusInput()

	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("==============program end %s!==============", diff)
}
