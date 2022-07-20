package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"gos/src/Micro/Protobuf/proto/userService"
)

func main() {
	u := &userService.UserInfo{
		Name:  "mark",
		Age:   10,
		Other: []string{"A", "B", "C"},
	}

	fmt.Println(u.GetOther())

	//protobuf序列化

	data, _ := proto.Marshal(u)
	fmt.Println(data)

	//反序列化
	var user userService.UserInfo
	_ = proto.Unmarshal(data, &user)
	fmt.Println(user)

}
