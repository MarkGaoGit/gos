package c4

import (
	"fmt"
	"strings"
)

func BBao() {
	sum := 0
	for i := 0; i <= 10; i++ {

		fc := func(i int) { sum += 1}


		fc(i)
	}
	fmt.Println(sum)

	fmt.Println()
	fmt.Println(NMing())

	fileSuffix := MakeFileSuffix("jpg")
	fmt.Println(fileSuffix("mp"))

}

// NMing 先return 1 然后进行匿名函数的运算 最后返回为2
func NMing() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}


func MakeFileSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + "." + suffix
		}
		return name
	}
}
