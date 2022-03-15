package Base

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const RootPath = "/Users/mark/work/go/gos"

var filePath = RootPath + "/storage/tmp.txt"

type Page struct {
	Title string
	Body  []byte
}

func FileReadWrite() {

	os.Stdout.WriteString("屏幕输出！ \n")

	//单行读取
	//eachReadLine()

	//读取所有
	//readAllFile()

	//按列读取
	//readColumnFile()

	//读取压缩包
	//readCompress()

	//写文件
	//pWriteFile()

	example()
}

func example() {
	ex := new(Page)
	ex.Title = "坏蛋"
	ex.Body = []byte("这是一本书")
	ex.save()
	content := ex.load("坏蛋")
	fmt.Println(content)
}

func (p *Page) save() {
	fp, err := os.OpenFile(RootPath+"/storage/exampleFile.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer fp.Close()

	handle := bufio.NewWriter(fp)
	handle.WriteString(p.Title + "-" + string(p.Body))
	handle.Flush()
}

func (p *Page) load(title string) string {
	fp, err := os.Open(RootPath + "/storage/exampleFile.txt")
	if err != nil {
		fmt.Println("文件打开失败")
		return ""
	}
	defer fp.Close()

	inputReader := bufio.NewReader(fp)
	for {
		inputString, readError := inputReader.ReadString('\n')
		if readError == io.EOF {
			return ""
		}
		lineContent := strings.Split(inputString, "-")
		if lineContent[0] == title {
			return lineContent[1]
		}
	}
}

func eachReadLine() {
	inputFile, inputError := os.Open(filePath)
	if inputError != nil {
		fmt.Println("文件读取失败" + inputError.Error())
		return
	}

	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readError := inputReader.ReadString('\n')
		if readError == io.EOF {
			fmt.Println("文件读取完成")
			return
		}
		fmt.Println(inputString)
	}
}

func readAllFile() {
	backFile := RootPath + "/storage/tmp_bak.txt"
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("文件读取失败")
		return
	}

	fmt.Println(string(buf))

	fmt.Println("文件内容写入备份文件!")

	err = ioutil.WriteFile(backFile, buf, 0777)
	if err != nil {
		fmt.Println("备份文件写入失败！")
	}

}

func readColumnFile() {
	fp, err := os.Open(RootPath + "/storage/couponData.csv")
	if err != nil {
		fmt.Println("文件读取失败！")
		return
	}
	defer fp.Close()

	var cl1, cl2, cl3 []string

	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(fp, &v1, &v2, &v3)
		if err != nil {
			fmt.Println("文件读取失败" + err.Error())
			break
		}
		cl1 = append(cl1, v1)
		cl2 = append(cl2, v2)
		cl3 = append(cl3, v3)
	}

	fmt.Println(cl1)
	fmt.Println(cl2)
	fmt.Println(cl3)

}

func readCompress() {
	var r *bufio.Reader
	fi, err := os.Open(RootPath + "/storage/test.gz")
	if err != nil {
		fmt.Println("文件读取失败" + err.Error())
		return
	}
	defer fi.Close()

	fz, err := gzip.NewReader(fi)
	if err != nil {
		r = bufio.NewReader(fi)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("文件读取完成！")
			return
		}
		fmt.Println("读取到内容" + line)
	}

}

func pWriteFile() {
	//读写权限打开文件 若文件不在则创建
	// 追加写文件
	fp, err := os.OpenFile(RootPath+"/storage/write.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件读取失败" + err.Error())
		return
	}
	defer fp.Close()

	handle := bufio.NewWriter(fp)

	for i := 0; i < 15; i++ {
		handle.WriteString("This is test file write" + strconv.Itoa(i) + " \n")
	}

	fmt.Fprint(handle, "文件写入完成\n")

	handle.Flush()

}
