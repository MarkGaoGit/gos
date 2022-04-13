package High

import (
	"fmt"
	"strconv"
)

type Person struct {
	Name    string
	money   float64
	changeF chan func()
}

// BfGetObject
//使用通道发送并访问对象
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.17.md
func BfGetObject() {
	bs := newPerson("mark", 19983949.5)
	fmt.Println(bs)
	bs.changeMoney(10000)
	fmt.Println(bs)

	fmt.Println("Salary changed:")
	fmt.Println(bs)
}

func newPerson(name string, money float64) *Person {
	p := &Person{name, money, make(chan func())}
	go p.executeP()
	return p
}

func (p *Person) changeMoney(newMoney float64) {
	p.changeF <- func() {
		p.money = newMoney
	}
}

func (p Person) getMoney() float64 {
	ch := make(chan float64)
	p.changeF <- func() {
		ch <- p.money
	}
	return <-ch
}

func (p *Person) executeP() {
	for f := range p.changeF {
		f()
	}
}

func (p *Person) String() string {
	return "name => " + p.Name + "有 " + strconv.FormatFloat(p.getMoney(), 'f', '2', 64) + "元"
}
