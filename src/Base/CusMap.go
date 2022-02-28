package Base

import (
	"fmt"
	"sort"
	"strconv"
)

func CusMap() {

	var m1 map[string]int
	var m2 map[string]int

	//map的两种初始化方式
	m1 = map[string]int{"mark": 30, "k2": 20}
	m3 := make(map[string]float32)

	m2 = m1
	m3["rankOne"] = 1.3
	m3["rankTow"] = 2.9
	m2["rankTmp"] = 100

	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)

	//容量为2个map 当多出来的时候会自动填充新的长度
	// 如果大概知道map的容量 最好是创建的时候添加容量
	group := make(map[string]string, 2)
	group["afterEnd"] = "this is after end group"
	group["beforeEnd"] = "this is before end group"
	fmt.Println(group)
	fmt.Println(len(group))
	group["middleEnd"] = "this is middle group"
	fmt.Println(len(group))

	//打印不存在的值的时候会输出值类型的空值
	fmt.Println(group["mark"])

	//多维的概念
	process := make(map[int][]int)
	process[9001] = []int{100, 938, 499, 1999}
	fmt.Println(process)
	for k, v := range process {
		if len(v) > 0 {
			for kk, vv := range v {
				fmt.Printf("pk => %d, kk => %d, vv => %d \n", k, kk, vv)
			}
		}
	}
	println()

	userInfo := make(map[string]string)
	userInfo["userName"] = "test"
	userInfo["group"] = "middlewareGroup"

	//输出fail 因为不存在权限的信息
	if _, ok := userInfo["permission"]; ok {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}

	userInfo["sex"] = "man"
	fmt.Println(userInfo)
	delete(userInfo, "sex")
	fmt.Println(userInfo)

	Week()

	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Println(items)

	MySort()

	kvChange()
}

func Week() {
	week := make(map[int]string, 7)
	for i := 1; i <= 7; i++ {
		week[i] = "周" + strconv.Itoa(i)
	}
	fmt.Println(week)
}

func MySort() {
	group := map[string]int{"c": 10, "a": 1, "z": 3, "f": 12, "d": 22}
	keys := make([]string, 5)

	i := 0
	for k, _ := range group {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for i := 0; i < len(keys); i++ {
		fmt.Println(group[keys[i]])
	}

}

func kvChange() {
	reveresMap := map[int]string{10: "mark", 100: "michael"}
	res := make(map[string]int, 3)

	for k, v := range reveresMap {
		res[v] = k
	}
	fmt.Println(res)
}
