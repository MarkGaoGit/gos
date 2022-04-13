package main

import (
	"fmt"
	"gos/src/High"
	"time"
)

func init() {
	fmt.Println("==============program start!==============")
}

func main() {
	start := time.Now()

	High.BfGetObject()

	end := time.Now()
	diff := end.Sub(start)
	fmt.Printf("==============program end %s!==============", diff)
}
