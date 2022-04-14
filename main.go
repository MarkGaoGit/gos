package main

import (
	Web "gos/src/Web/Server"
)

func init() {
	//fmt.Println("==============program start!==============")
}

func main() {
	//start := time.Now()

	Web.HttpServer()

	//end := time.Now()
	//diff := end.Sub(start)
	//fmt.Printf("==============program end %s!==============", diff)
}
