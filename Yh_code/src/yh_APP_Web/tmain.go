package main

import (
	"fmt"
	"io"
	"net/http"
)

// 初始化
func init() {
	fmt.Println("init func!!!")

	return
}

//  主函数
func main() {
	http.HandleFunc("/test", yh_web_server)
	http.ListenAndServe(":7878", nil)
	return
}

// 处理函数
func yh_web_server(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello 航语web!!!")
	return
}
