package Base

import (
	"fmt"
	"io"
	"os"
)

func CopyFile() {

	//文件拷贝
	fileHandle(filePath, RootPath+"/storage/tmp_copy.txt")

}

func fileHandle(source, target string) {
	sourceHd, err := os.Open(source)
	if err != nil {
		fmt.Println("文件打开失败！")
		return
	}
	defer sourceHd.Close()

	targetHd, err := os.Create(target)
	if err != nil {
		fmt.Println("文件创建失败！")
		return
	}
	defer targetHd.Close()

	_, err = io.Copy(targetHd, sourceHd)
	if err != nil {
		fmt.Println("文件拷贝失败！")
		return
	}
	fmt.Println("文件拷贝完成！")

}
