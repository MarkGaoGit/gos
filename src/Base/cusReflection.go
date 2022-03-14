package Base

import (
	"fmt"
	"reflect"
)

type TagType struct {
	field1 bool  "这是一个bool值"
	field  int32 "这是一个int值"
}

type IDuck interface {
	Quack()
	Walk()
}

type Bird struct {
	name string
}

func (b *Bird) Quack() {
	fmt.Println("quack executive")
}

func (b *Bird) Walk() {
	fmt.Println("walk executive")
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 2; i++ {
		duck.Quack()
		duck.Walk()
	}
}

func CusReflection() {

	tt := TagType{true, 33}

	for i := 0; i < 2; i++ {
		refTag(tt, i)
	}

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type())
	fmt.Println("settability of v:", v.CanSet())
	v = v.Elem()
	fmt.Println("The Elem of v is: ", v)
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface())

	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()

	bird := new(Bird)
	DuckDance(bird)

}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Println(ttType)
	fmt.Println(ixField)
}
