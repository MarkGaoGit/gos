package Base

import "fmt"

const Pi = 3.14159
const Key string = "\ncache:Token"

const Ln2 = 0.1913843298342432
const Log2 = 1 / Ln2
const Billion = 1e9

// 可以用iota
const (
	Unknown = 0
	Man     = 1
	Woman   = 2
)

const (
	Start   = iota //0
	Ding           //1
	Success        //2
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
)

func FcConst() {
	fmt.Print(Pi)
	fmt.Println(Key)

	fmt.Print(Ln2)
	fmt.Println("\n")

	fmt.Print(Log2)
	fmt.Println("\n")

	fmt.Print(Billion)
	fmt.Println("\n")

	fmt.Println(Unknown + Man + Woman)
	fmt.Println("\n")

	fmt.Println(Start)
	fmt.Println("\n")
	fmt.Println(Ding)
	fmt.Println("\n")
	fmt.Println(Success)
	fmt.Println("\n")

	fmt.Print(KB)
	fmt.Println("\n")
	fmt.Print(MB)
	fmt.Println("\n")
	fmt.Print(GB)
	fmt.Println("\n")
}
