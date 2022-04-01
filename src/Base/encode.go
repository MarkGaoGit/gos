package Base

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func EncodeTmp() {
	hashShar1 := sha1.New()
	io.WriteString(hashShar1, "test")
	var b []byte
	sharResult := hashShar1.Sum(b)
	fmt.Printf("shar1：%x \n", sharResult)

	hashShar1.Reset()

	data := []byte("123456")
	n, err := hashShar1.Write(data)
	if err != nil || n != len(data) {
		log.Printf("hash error")
	}

	checkSum := hashShar1.Sum(b)
	fmt.Printf("shar1: %x\n", checkSum)

	// MD5
	md5T := md5.New()
	io.WriteString(md5T, "123456")
	var a []byte
	fmt.Printf("MD5：%x \n", md5T.Sum(a))

}
