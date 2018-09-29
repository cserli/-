package main

import (
	"fmt"
	"net/http"
)

// 初始化操作
func init() {
	fmt.Println("航语网关")
	return
}

// 微服务框架 网关服务器
func main() {
	http.HandleFunc("/Auth", AuthFunc)
	http.ListenAndServe(":8888", nil)
	return
}

// 数据转发到 grpc 认证服务器操作
func AuthFunc(w http.ResponseWriter, r *http.Request) {
	// 数据发送到grpc 认证服务器
	io.WriteString(w, "hello 航语认证!")
	return
}
