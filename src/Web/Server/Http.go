package Web

import (
	"fmt"
	"io"
	"net/http"
)

// HttpServer
// 启动http服务
//https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.2.md
func HttpServer() {
	//2种方式都可以访问到具体的方法
	http.HandleFunc("/create", defaultRoute)

	//http.Handle("/create", http.HandlerFunc(defaultRoute) )

	// 若启动https http.ListenAndServeTLS
	err := http.ListenAndServe("localhost:2000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func defaultRoute(w http.ResponseWriter, req *http.Request) {
	a := req.FormValue("a")

	io.WriteString(w, "这是第一种写法\n")
	fmt.Fprintf(w, "requestPath："+req.URL.Path[1:]+"\n")
	fmt.Fprintf(w, "requestParamsA："+a+"\n")
}
