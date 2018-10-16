package main

import (
	//	"crypto/aes"
	//	"crypto/cipher"
	//	"crypto/rand"
	"encoding/json"
	"fmt"
	"glog-master"

	//	"io"
	"log"
	"net/http"
	"os"

	//	"strings"
	//	"bytes"
	"encoding/base64"
	"time"

	//	"unicode/utf8"

	"code.google.com/p/mahonia"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld" // 数据处理  protobuf 内服务
)

// 47.104.209.25
// 七牛云的公网IP
// 对象存储操作
// 视频直播
// 数据操作等

const (
	address     = "localhost:50051"
	defaultName = "world"
)

type SSSS struct {
	Name string
}

// 数据存储 --
type SSSSbak struct {
	Data interface{}
}

// 数据处理操作
func WenDaOrTuCao(strnickName, stravatarUrl, strdata string, w http.ResponseWriter) {
	glog.Info("strnickName, stravatarUrl, strparam", strnickName, stravatarUrl, strdata)
	// 服务地址
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//  客户端申请操作 --  也就是相当于注册一个客户端事件
	c := pb.NewGreeterClient(conn)
	name := strnickName + "☢" + stravatarUrl + "☢" + strdata
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second) // 超时设置 --
	defer cancel()
	// 远程回调  ---  函数的实现操作的
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
	data := SSSS{
		Name: r.Message,
	}

	b, _ := json.Marshal(data)
	log.Printf("Greeting b: %s", string(b))
	// fmt.Fprint(w, r.Message)
	// str := ConvertToString(string(b), "gbk", "utf-8")
	// bbw := PswEncrypt(string(b))
	// bbw, _ := Encrypt(b, []byte(sKey))
	// bbw := PswEncrypt("abcsfsssfsfsfsfsfsfsfsfs3535353535353533")
	// base64
	// bbbb := bbw // + "\x00"
	encodeString := base64.StdEncoding.EncodeToString(b)
	ddd := SSSSbak{
		Data: encodeString,
	}

	bb, _ := json.Marshal(ddd)
	log.Printf("Greeting b: %s", string(bb))
	// fmt.Fprint(w, string(bb))
	fmt.Fprint(w, string(bb))
	return
}

// 数据处理
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
