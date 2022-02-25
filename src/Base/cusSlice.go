package Base

import (
	"fmt"
	"sort"
)

// CusSlice 切片可以理解为是数组的全部 或者是数组的一部分
func CusSlice() {

	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			screen[x][y] = 1
		}
	}

	// 0 <= len(s) <= cap(s)

	s := screen[0:3]

	fmt.Println(s)

	//切片真是存在的数据大小
	fmt.Println(len(s))
	//切片的容量
	fmt.Println(cap(s))

	x := []int{2, 3, 5, 7, 11}

	/**
	x是一个数组
	当 x[1:3]时 表示 从x的第一个元素开始到最后 为切片y 这时 len(y)= 2 cap(y) = 4
	也就是说从第1位开始 截取了后面所有的元素为一个切片
	那么y[0:4] 就是 3， 5， 7， 11
	当 y[0:5]超出切片长度 则错误
	*/
	y := x[1:3]
	fmt.Println(y)      // 3, 5
	fmt.Println(y[0:4]) //  3, 5, 7, 11

	arr := [5]int{0, 1, 2, 3, 4}
	CSum(arr[:])

	s1 := make([]int, 5)
	fmt.Println(s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	//切片长度是5 容量是10
	s2 := make([]string, 5, 10)
	fmt.Println(s2)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	//new 和 make区别 就是 new是分配了内存 返回的是一个指针 而make则是返回一个类型的初始值
	a := new([]int)
	b := make([]int, 0)
	fmt.Println(a)
	fmt.Println(b)

	// 切片是从数组中截取的相当于获取的指针， 那么切片中的值变更后 数组中的值也是会变更的
	// 简单的理解就是 引用
	s3 := []byte{'p', 'o', 'e', 'm'}
	s4 := s3[2:]
	fmt.Println(s4) //e m
	s4[1] = 't'
	fmt.Println(s3) //p o e t
	fmt.Println(s4) //e t

	items := [...]int{10, 20, 30, 40}
	for k, v := range items {
		items[k] = v * 2
	}
	fmt.Println(items)

	var sumAr int
	var avgAr float32

	var tmpArr = []int{1, 2, 34, 6}
	sumAr, avgAr = CArraySum(tmpArr[:])
	fmt.Println(sumAr, avgAr)

	ar := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(ar)
	fmt.Println(ar[3:3]) //没有截取 空的
	fmt.Println(ar[3:4]) // 截取到了4

	//++++++++++++++++++++++++++++++切片复制+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	form := []int{1, 2, 3, 4, 5}
	to := make([]int, 5)

	//当填充的切片容量不足时将会截断后续的值
	copy(to, form)
	fmt.Println(to)

	//追加会分配新的内存
	to = append(to, 6, 7, 8)
	fmt.Println(to)

	//++++++++++++++字符串、数组、切片应用+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	tmpStr := "this is string"
	slice := []byte(tmpStr)
	var slice1 = make([]byte, 10)
	fmt.Println(slice)
	copy(slice1, tmpStr)
	fmt.Println(slice1)
	fmt.Println(tmpStr[:3])

	//字符串变更必须先 把字符串转换为数组
	slice[0] = 'T'
	tmpStr2 := string(slice)
	fmt.Println(tmpStr2)

	sortArr := []int{10, 3, 10, 22, 33, 1, 2, 3, 6, 5, 4}
	sort.Ints(sortArr)
	isSort := sort.IntsAreSorted(sortArr)
	fmt.Println(isSort)
	fmt.Println(sortArr)

	//排序以后此才可以使用的查询，应为查询是用的二分法
	fmt.Println(sort.SearchInts(sortArr, 10))

	slice10 := []int{1, 2, 4, 5, 6, 7}
	fmt.Println(append(slice10[:2], slice10[3:]...)) //删除3
	fmt.Println(append(slice10, make([]int, 5)...))  //扩展5个容量
	//a = append(a[:i], append([]T{x}, a[i:]...)...)
	fmt.Println(append(slice10[:2], append([]int{3}, slice10[2:]...)...)) //在索引2的地方加一个值

	r1, r2 := StrCup("This is testing", 3)
	fmt.Println(r1, r2)

	//go 实现的冒泡
	Mp()
}

func Mp() {
	arr := []int{10, 2, 3, 1, 9, 5, 4}

	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[i] > arr[j] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
	}
	fmt.Println(arr)

}

func StrCup(s string, i int) (r1 string, r2 string) {
	return s[:i], s[i:]
}

// CArraySum 返回一个切片的和 和 平均值
func CArraySum(arrF []int) (sumAr int, avgAr float32) {
	for _, v := range arrF {
		sumAr += v
	}

	return sumAr, float32(sumAr / len(arrF))
}

func CSum(a []int) {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}
