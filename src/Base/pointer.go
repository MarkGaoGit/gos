package Base

import "fmt"

//Pointer 指针
func Pointer() {

	var number = 10
	fmt.Printf("内存地址%p\n", &number)

	var intP *int
	intP = &number
	fmt.Printf("内存地址%p, 引用的值%d", intP, *intP)

}
