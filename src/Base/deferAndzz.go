package Base

import "fmt"

// GetOne 测试defer
func GetOne() {
	ConnectDB()
	defer CloseDB()

	sql := "SELECT * FROM m_sms_template"
	ExecuteSQL(sql)
}

func ConnectDB() {
	fmt.Println("DB connect success!")
}

func CloseDB() {
	fmt.Println("DB close success!")
}

func ExecuteSQL(sql string) {
	fmt.Println(sql)
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func Aa() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func Bb() {
	defer un(trace("b"))
	fmt.Println("in b")
	Aa()
}
