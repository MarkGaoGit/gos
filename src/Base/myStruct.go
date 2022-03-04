package Base

import (
	"fmt"
	"gos/src/MPack"
	"strings"
)

type userInfo struct {
	name  string
	age   int
	desc  string
	money float64
}

type rectangleWH struct{ x, y int }

type File struct {
	fd   int
	name string
}

//匿名结构体
type inS struct {
	in1 int
	in2 int
}

type outS struct {
	a int
	b int
	float32
	string
	inS
}

func MyStruct() {
	user := new(userInfo)
	user.name = "mark"
	user.desc = `This is user`
	user.age = 20
	user.money = 10092.12345234
	fmt.Println(*user)

	var michael userInfo
	michael.name = "michael"
	NameToUp(&michael)
	fmt.Println(michael)

	//矩形面积
	rectangle := rectangleWH{3, 4}
	fmt.Println(rectangle.rectangleArte())

	//文件处理工厂方法
	fmt.Println(*NewFile(10, "./package.go"))

	//自定义包中的结构体
	packageStruct := new(MPack.TmpStruct)

	packageStruct.X = 1
	packageStruct.Y = 10
	fmt.Println(packageStruct)

	// 匿名结构体
	outer := new(outS)
	outer.a = 10
	outer.b = 20
	outer.float32 = 1000.23
	outer.in1 = 30
	outer.in2 = 40
	fmt.Println(*outer)

	outer2 := outS{100, 200, 300., "mark", inS{400, 500}}
	fmt.Println(outer2)
}

func NameToUp(u *userInfo) {
	u.name = strings.ToUpper(u.name)
}

// 计算矩形面积 传入条件 x y 的值 结构函数！
func (c rectangleWH) rectangleArte() int {
	return c.x * c.y
}

// NewFile 文件处理的工厂方法
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}
