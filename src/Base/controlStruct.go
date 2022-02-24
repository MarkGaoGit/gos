package Base

import (
	"fmt"
	"runtime"
	"strings"
)

func ControlStruct() {
	age := 30

	if age < 10 {
		fmt.Print("未成年")
	} else if age > 10 && age < 50 {
		fmt.Println("中年")
	} else {
		fmt.Println("老年")
	}

	if runtime.GOOS == "windows" {
		fmt.Println("windows")
	} else {
		fmt.Println(runtime.GOOS)
	}

	switch 3 % 2 {
	case 1:
		fmt.Println("奇数")
	case 0:
		fmt.Println("偶数")
	}

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	fmt.Println()

	for i := 0; i <= 10; i++ {
		if i%2 < 1 {
			continue
		} else {
			fmt.Println(i)
		}
	}

	for i := 0; i <= 5; i++ {
		var tmp string
		tmp = strings.Repeat("G", i)
		fmt.Printf("%s \n", tmp)
	}

	tmpStr := "this is testing"
	for pos, val := range tmpStr {
		fmt.Printf("k => %c, v => %d \n", val, pos)
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

	//kk := 0
	//POINT:
	//	print(kk)
	//	kk++
	//
	//if kk == 5 {
	//	return
	//}
	//goto POINT

	iof := 0
	for {

		if iof >= 3 {
			break
		}
		//break out of this for loop when this condition is met
		fmt.Println("Value of i is:", iof)
		iof++
	}
	fmt.Println("A statement just after for loop.")

	for i := 0; i < 7; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Odd:", i)
	}

}
