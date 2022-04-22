package Web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Dove struct {
	AppName       string `json:"appName"`
	Environment   string `json:"environment"`
	AppLogPath    string `json:"appLogPath"`
	PhpVersion    string `json:"phpVersion"`
	SwooleVersion string `json:"swooleVersion"`
	Datetime      string `json:"datetime"`
}

// RequestAndReadPage
// 请求 和 获取请求的页面
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.3.md
func RequestAndReadPage() {
	urlPath := "https://dove.rp-field.com"

	content, err := http.Get(urlPath)
	checkError(err)

	data, err := ioutil.ReadAll(content.Body)

	res := &Dove{}
	json.Unmarshal(data, &res)

	fmt.Println(res)

}
