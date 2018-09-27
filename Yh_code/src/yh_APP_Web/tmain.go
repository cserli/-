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
	http.HandleFunc("/", helloWorldHandler)
	http.ListenAndServe(":80", nil)
}

// 处理函数
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello 航语web!!!")
	return
}
