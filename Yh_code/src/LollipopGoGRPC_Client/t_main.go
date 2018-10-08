package main

import (
	"fmt"
	"glog-master"
	"net/http"
)

// 接受数据处理
func TJWanJiaData(w http.ResponseWriter, req *http.Request) {
	glog.Info("httpTask is running...")
	if req.Method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.ParseForm()
		if true {
			Protocol, bProtocol := req.Form["variable"]
			glog.Info("httpTask is running...", Protocol, bProtocol)
			//PswDecrypt(Protocol[0])
			WenDaOrTuCao("test", "test2", "test3", w)
		}
		return
		// 获取函数
		Protocol, bProtocol := req.Form["Protocol"]
		Protocol2, bProtocol2 := req.Form["Protocol2"]
		glog.Info("httpTask is running...", Protocol, bProtocol, Protocol2, bProtocol2)
		fmt.Fprint(w, "数据接受成功！！！")
		return
		if bProtocol && bProtocol2 {
			// 主协议判断
			if Protocol[0] == "1" {
				switch Protocol2[0] {
				case "2": // 发表吐槽
					{
						strnickName, _ := req.Form["nickName"]
						stravatarUrl, _ := req.Form["avatarUrl"]
						strparam, _ := req.Form["param"]
						glog.Info("strparam", strparam)
						// 发送给 gRPC--server
						WenDaOrTuCao(strnickName[0], stravatarUrl[0], strparam[0], w)
						break
					}
				default:
					fmt.Fprint(w, "server Protocol2 default is Error！！！")
					return
				}
			}
		}
	}
}

// 主函数  http + grpc 高性能的数据处理操作
func main() {
	http.HandleFunc("/test", TJWanJiaData) // 测试操作
	// http.HandleFunc("/Auth", AuthFunc)     // 认证服务器  --- 微信+短信认证操作
	err := http.ListenAndServe(":7878", nil)
	//err := http.ListenAndServeTLS(":7878", "cert.pem", "key.pem", nil)
	if err != nil {
		glog.Info("Entry nil", err.Error())
		return
	}
}
