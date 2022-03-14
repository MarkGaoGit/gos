package Base

import (
	"container/list"
	"fmt"
	"gos/src/MPack"
	"regexp"
	"strconv"
	"unsafe"
)

// 小写字母开头的函数 只作用于包内
// 大写字母开头的函数 包外也是可见的

func MyPackage() {

	MyList()

	byteSize := 10
	fmt.Println(unsafe.Sizeof(byteSize))

	MyRegexpMatch("This is test string 1000.3")
	fmt.Println(0.1 + 0.2)

	//自定义包
	fmt.Println(MPack.MDay("mark"))
	fmt.Println(MPack.MNight("mark"))
	fmt.Println(MPack.IsAm())
	fmt.Println(MPack.IsPm())

}

// MyList 双向链表
func MyList() {
	mList := list.New()
	mList.PushBack(100)
	mList.PushFront(101)
	mList.PushBack(102)

	for e := mList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// MyRegexpMatch 正则匹配并替换
func MyRegexpMatch(str string) {

	pat := "\\d.+."

	f := func(s string) string {
		fmt.Println(s)
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(str)); ok {
		fmt.Println("匹配到了")
	} else {
		fmt.Println("没有匹配到么！")
	}

	res, _ := regexp.Compile(pat)

	tmpStr := res.ReplaceAllString(str, "check regexp")
	fmt.Println(tmpStr)
	tmpStrFunc := res.ReplaceAllStringFunc(str, f)
	fmt.Println(tmpStrFunc)

}
