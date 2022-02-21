package c4

import (
	"fmt"
	"strconv"
	"strings"
)

type Str string

// StringS 字符串
func StringS() {
	var description, name Str = `This is a coupon description`, `mark`

	var tmpStr string = "this is testing"

	fmt.Println(description)
	fmt.Println(name)

	// 使用 + 或者 += 来连接2个字符串
	fmt.Println(description + ` ` + name)

	// 字符串长度
	descriptionStrLen := len(description)
	fmt.Println(descriptionStrLen)

	// 是否以某字符串开头
	fmt.Printf("%t\n", strings.HasPrefix(tmpStr, "T"))
	fmt.Printf("%t\n", strings.HasSuffix(tmpStr, "ing"))

	fmt.Println(strings.Index(tmpStr, "t"))

	b := strings.Replace(tmpStr, "t", "T", -1)
	fmt.Println(b)

	//slice中字符串类型的元素 分割
	slices := []string{"a", "b", "c"}
	fmt.Println(strings.Join(slices, " + "))

	fmt.Println(strings.Split(tmpStr, "="))

	//转换为字符串
	age := 33
	fmt.Println(strconv.Itoa(age))

}




