package main

import "gos/src/Web"

func init() {
	//fmt.Println("==============program start!==============")
}

func main() {
	//start := time.Now()

	go Web.RPCServer()
	Web.RPCClient()

	//end := time.Now()
	//diff := end.Sub(start)
	//fmt.Printf("==============program end %s!==============", diff)
}
