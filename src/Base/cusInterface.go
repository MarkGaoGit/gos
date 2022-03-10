package Base

import (
	"fmt"
	"sort"
)

// Method 若接口中只包含一个方法 那么名字需要使用er结尾
type Method interface {
	WorkStart() string
	Sleep(ttl int)
}

type Personal struct {
	name string
	age  int
	Method
}

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInfo(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

type Element interface{}

func CusInterface() {
	per := new(Personal)
	per.name = "mark"
	start := per.WorkStart()
	fmt.Println(start)
	per.Sleep(10)

	Classifier(13, 33, 22.5, nil, true, "this is test string")

	var lst List
	if LongEnough(lst) {
		fmt.Printf("xxxxxx")
	}

	pslt := new(List)
	CountInfo(pslt, 1, 100)
	if LongEnough(pslt) {
		fmt.Printf("llllllll \n")
	}
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntSlice(data[0:])
	SortInterfaceMp(a)

	var ele Element
	ele = []int{10, 20}
	fmt.Println(ele)

}

// SortInterfaceMp SortCus 传入接口类型 然后冒泡排序
func SortInterfaceMp(data sort.Interface) {
	for i := 0; i < data.Len(); i++ {
		for j := 1; j < data.Len()-1; j++ {
			if data.Less(i, j) {
				data.Swap(i, j)
			}
		}
	}
	fmt.Println(data)
}

func Classifier(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case int, int8, int32:
			fmt.Printf("int %d\n", v)
		case bool:
			fmt.Printf("bool %d \n", v)
		case string:
			fmt.Printf("string %s \n", v)
		case float64, float32:
			fmt.Printf("float %d \n", v)
		default:
			fmt.Printf("not v %s \n", v)
		}
	}
}

func (p *Personal) WorkStart() string {
	return p.name + " 开始工作"
}

func (p *Personal) Sleep(ttl int) {
	fmt.Printf("休息了%d个小时 \n", ttl)
}
