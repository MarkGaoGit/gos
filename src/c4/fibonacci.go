package c4

import (
	"fmt"
)

const LIM = 10

var fibs [LIM]uint64

func Fibonacci() {
	var result uint64 = 0

	for i := 0; i < LIM; i++ {
		result = FibonacciBody(i)
		fmt.Printf("number %d is %d \n", i, result)
	}

}

func FibonacciBody(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = FibonacciBody(n - 1) + FibonacciBody(n - 2)
	}
	fibs[n] = res
	return
}





