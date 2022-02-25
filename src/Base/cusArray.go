package Base

import "fmt"

const (
	WIDTH  = 10
	HEIGHT = 5
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func CusArray() {
	//var source [5]int

	//定长数组
	source := [5]int{}

	source[3] = 100

	fmt.Println(len(source))

	for k, v := range source {
		fmt.Printf("K => %d V => %d \n", k, v)
	}

	//不定长数组
	array2 := [...]int{1, 2, 3, 4, 5}
	for k, v := range array2 {
		fmt.Printf("K => %d V => %d \n", k, v)
	}

	array3 := []string{3: "Michael", 5: "Mario"}
	fmt.Println(array3)

	//传入指针
	for i := 0; i <= 3; i++ {
		Fp(&[3]int{i, i * i, i * i * i})
	}

	//矩形数组
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 3
		}
	}

	// 如果是大型数组则传入地址 函数中使用指针接收参数
	// 函数中的值变更后 传入的值也会变
	Fp2(&screen)
	fmt.Println(screen)

}

func Fp(p *[3]int) {
	fmt.Println(p)
}

func Fp2(p *[WIDTH][HEIGHT]pixel) {
	p[WIDTH-1][HEIGHT-1] = 1
	fmt.Println(p)
}
