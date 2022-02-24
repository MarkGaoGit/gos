package Base

import (
	"fmt"
	"math"
	"math/rand"
)

var n int16 = 40
var m int32

// Ys 运算符
func Ys() {
	m = int32(n)
	print(m)

	b := 1.1
	c := 2
	b++
	c++
	println("\r")
	print(c)
	println("\r")
	print(b)

}

func Uint8FromIn(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d 超出了int8的值范围", n)
}

// RandNumber 随机数
func RandNumber(n int) int {
	a := 0
	for i := 0; i < n; i++ {
		a := rand.Int()
		fmt.Printf("%d", a)
		println("\r")
	}
	return a
}

type StringStr string

// TypeAlias 类型别名
func TypeAlias() {
	var userName StringStr = "李四"
	println(userName)
}
