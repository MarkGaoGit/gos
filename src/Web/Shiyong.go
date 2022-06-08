package Web

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type TmpInterface interface {
	a()
	b()
	c()
}

func ShiyongFunc() {
	strUpdate()

	sliceArray()

	throwException(func() {
		panic("这是抛了一个错误")
	})

	fileRead()

}

func strUpdate() {
	str := "hello"
	str1 := []byte(str)
	str1[2] = 'R'
	fmt.Println(string(str1))

	//从第二个截取到第三个
	fmt.Println(str[2:3])

	//字符串拼接
	str2 := "mi-mi "
	str2 += string(str1)
	fmt.Println(str2)
}

func sliceArray() {
	hasTrue := false
	array := [...][]string{{"a", "b", "c"}, {"A", "B", "C"}}
	pointerUse := "c"

Found:
	for _, item := range array {
		for _, v := range item {
			if v == pointerUse {
				fmt.Println(v)
				hasTrue = true
				break Found
			}
		}
	}

	fmt.Println(hasTrue)

	//查看一个数组是否存在设置
	users := map[string]string{"name": "michael", "age": "22"}

	if _, ok := users["ccc"]; ok {
		fmt.Println(users["ccc"])
	} else {
		fmt.Println("没有找到数组的这个下标")
	}

}

func throwException(g func()) {
	defer func() {
		if x := recover(); x != nil {
			log.Printf("run time panic %v", x)
		}
		log.Print("done")
	}()

	log.Println("start")
	g()
}

func fileRead() {
	file, err := os.Open("./storage/tmp.txt")
	if err != nil {
		fmt.Println("文件读取失败！")
		return
	}
	defer file.Close()
	iReader := bufio.NewReader(file)
	for true {
		str, err := iReader.ReadString('\n')
		if err != nil {
			return
		}
		fmt.Printf("The input was:【 %s 】", str)
	}
}
