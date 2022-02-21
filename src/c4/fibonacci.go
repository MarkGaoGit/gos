package c4

import (
	"fmt"
)

func Fibonacci() {
	result := 0

	for i := 0; i < 3; i++ {
		result = FibonacciBody(i)
		fmt.Printf("number %d is %d \n", i, result)
	}

}

func FibonacciBody(n int) (res int) {
	if n <= 1 {
		return 1
	} else {
		return FibonacciBody(n - 1) + FibonacciBody(n - 2)
	}
}
