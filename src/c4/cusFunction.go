package c4

import "fmt"

var num, nums = 10, 20

var num1, num2, sumN, minN, cN int

// CusFunction 自定义函数
func CusFunction() {

	num1, num2 = Xnum(num)
	fmt.Printf("num = %d, num1 = %d, num2 = %d \n", num, num1, num2)

	num1, num2 = Xnum2(num)
	fmt.Printf("num = %d, num1 = %d, num2 = %d \n", num, num1, num2)

	sumN = SumNumber(num, nums)
	fmt.Printf("he = %d \n", sumN)


	minN = JianNumber(num, nums)
	fmt.Printf("jian = %d \n", minN)

	cN = CNumber(num, nums)
	fmt.Printf("cheng = %d \n", cN)

	n := 0
	reply := &n
	CusPointer(10, 5, reply)
	print(*reply)

	x := Min(1, 2, 3, 0, 10, -1)
	fmt.Printf("this is minnumber %d \n", x)

	s := []int{10, 20, 2, 30, -2, -1, 1000}
	minNum := Min(s...)
	fmt.Printf("this is minnumber by slice %d \n", minNum)

}

func Xnum(input int) (int, int) {
	return 2 * input, 3 * input
}

func Xnum2(input int) (x2, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

func SumNumber(p1, p2 int) (he int) {
	return p2 + p1
}

func JianNumber(p1, p2 int) (ji int) {
	return p2 - p1
}


func CNumber(p1, p2 int) (cNumber int) {
	cNumber = p1 * p2
	return
}

func CusPointer(a, b int, reply *int) {
	*reply = a * b
}

func Min(s ...int) int {
	min := s[0]

	for _, v := range s {
		fmt.Println(v)
		if v < min {
			min = v
		}
	}

	return min
}