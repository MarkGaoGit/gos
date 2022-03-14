package Base

import (
	"bufio"
	"fmt"
	"os"
)

var (
	firstName, lastName, s string
	i, age                 int
)

func ReadCusInput() {

	//读取用户输入的姓名
	//fmt.Println("place input name")
	//fmt.Scanln(&firstName, &lastName)
	//fmt.Println(firstName, lastName)

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("输入内容：")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("读取输入内容失败！")
		return
	}

	fmt.Println("输入的是" + input)

	switch input {
	case "Philip\n":
		fmt.Println("Welcome Philip!\n")
	case "Chris\n":
		fmt.Println("Welcome Chris!\n")
	case "Ivo\n":
		fmt.Println("Welcome Ivo!\n")
	default:
		fmt.Printf("V1 You are not welcome here! Goodbye!\n")
	}

	// version 2: 使用fallthrough 就很nice了 后面的都会执行的
	switch input {
	case "Philip\n":
		fallthrough
	case "Ivo\n":
		fallthrough
	case "Chris\n":
		fmt.Printf("Welddddcome %s\n", input)
	default:
		fmt.Printf("V2 You are not welcome here! Goodbye!\n")
	}

	switch input {
	case "Philip\n", "Ivo\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Printf("V3 You are not welcome here! Goodbye!\n")
	}
}

func ReadExample1() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('S')
	//fmt.Println(input)
	//fmt.Println(len(input))

	for _, m := range input {
		fmt.Println(string(m))
	}
}
