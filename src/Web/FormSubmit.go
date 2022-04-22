package Web

import (
	"fmt"
	"io"
	"net/http"
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

func simpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello word</h1>")
}

func save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, r.FormValue("in"))

	}
}

// FormServer https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/15.4.md
func FormServer() {
	http.HandleFunc("/test", simpleServer)
	http.HandleFunc("/save", save)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		fmt.Println("服务启动失败！")
	}
	fmt.Println("server staring....")
}
