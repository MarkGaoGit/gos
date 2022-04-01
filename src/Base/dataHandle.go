package Base

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

var jsonFile = RootPath + "/storage/tmp.json"

type ShoesColor struct {
	ColorCode string
}

type Shoes struct {
	Brand  string
	Size   int
	Colors []ShoesColor
	Gender string
}

func DataHandle() {
	encodeJson()
	decodeJson()
	getFileData()

	decodeXml()
}

func decodeXml() {
	var t xml.Token
	var err error
	input := "<Person><FirstName>Laura</FirstName><LastName>Lynn</LastName></Person>"
	inputReader := strings.NewReader(input)
	p := xml.NewDecoder(inputReader)
	for t, err = p.Token(); err == nil; t, err = p.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Printf("Token name: %s\n", name)
			for _, attr := range token.Attr {
				attrName := attr.Name.Local
				attrValue := attr.Value
				fmt.Printf("An attribute is: %s %s\n", attrName, attrValue)
				// ...
			}
		case xml.EndElement:
			fmt.Println("End of token")
		case xml.CharData:
			content := string([]byte(token))
			fmt.Printf("This is the content: %v\n", content)
			// ...
		default:
			// ...
		}
	}

}

func decodeJson() {

	//不知道返回格式的时候可以直接适用interface 来接收decode的值
	var data interface{}
	//var data map[string]string
	encodeData := []byte(`{
    "appName": "zlphp-backend-middleware",
    "environment": "pre",
    "appLogPath": "/opt/www/runtime",
    "phpVersion": "7.4.21",
    "swooleVersion": "4.6.0",
    "datetime": "2022-04-01 11:24:02"
}`)
	json.Unmarshal(encodeData, &data)

	//for k, v := range data {
	//	fmt.Printf("k => %s, v=> %s \n", k, v)
	//}

}

func encodeJson() {

	// JSON解析时结构体属性的首字母必须大写
	scr := ShoesColor{"red"}
	scy := ShoesColor{"yellow"}
	shoesInfo := Shoes{"Nike", 44, []ShoesColor{scr, scy}, "man"}
	shoesInfoJson, _ := json.Marshal(shoesInfo)
	fmt.Printf("JsonEncode： %s\n", shoesInfoJson)

	//saveFile(shoesInfoJson)
}

// json数据保存到文件中
func saveFile(jsonData []byte) {
	fp, _ := os.OpenFile(jsonFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer fp.Close()
	jFp := json.NewEncoder(fp)
	jFp.Encode(jsonData)
}

func getFileData() {
	var sh Shoes
	fp, _ := os.Open(jsonFile)
	defer fp.Close()
	jFp := json.NewDecoder(fp)
	jFp.Decode(sh)
	fmt.Println(sh)
}
