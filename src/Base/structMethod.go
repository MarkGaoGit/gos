package Base

import (
	"fmt"
	"strconv"
)

//三角形信息
type triangleInfo struct {
	height int
	bottom int
	edge1  int
	edge2  int
}

type Engine interface {
	Start()
	End()
}

type Che struct {
	Engine
	wheelCount int
}

type Mercedes struct {
	Che
}

type TowInts struct {
	a int
	b int
}

func StructMethod() {
	triangle := &triangleInfo{100, 30, 40, 60}
	triangleArea := triangle.areaOrPerimeter(true)
	trianglePerimeter := triangle.areaOrPerimeter(false)
	fmt.Println(triangleArea, trianglePerimeter)

	b := new(Mercedes)
	b.wheelCount = 10
	b.sayHiToMerkel()

	tmp := new(TowInts)
	tmp.a = 100
	tmp.b = 200
	fmt.Printf("two1 is: %v\n", tmp)
	fmt.Println("two1 is:", tmp)
	fmt.Printf("two1 is: %T\n", tmp)
	fmt.Printf("two1 is: %#v\n", tmp)

}

func (t *TowInts) String() string {
	return "(" + strconv.Itoa(t.a) + "/" + strconv.Itoa(t.b) + ")"
}

func (c *Mercedes) sayHiToMerkel() {
	fmt.Println("Engine init")
	c.Start()
	c.End()
	c.numberOfWheels()
}

func (c *Che) GoToWork() {
	c.Start()
	c.End()
}

func (c *Che) numberOfWheels() int {
	return c.wheelCount
}

func (c Che) Start() {
	fmt.Println("Che Start")
}

func (c *Che) End() {
	fmt.Println("Che End")
}

// Height 设置和获取三角形的基础信息
func (t *triangleInfo) Height() int {
	return t.height
}

func (t triangleInfo) SetHeight(h int) {
	t.height = h
}

// 三角形面积或周长计算
// 若 triangleInfo 为 *triangleInfo 则函数中可以改变 三角形的基础信息
// 传入指针更省内存
func (t *triangleInfo) areaOrPerimeter(isArea bool) float64 {
	if isArea {
		return float64(t.height * t.bottom / 2)
	}
	return float64(t.bottom + t.edge1 + t.edge2)
}
