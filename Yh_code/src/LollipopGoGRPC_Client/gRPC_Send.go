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

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

// 问答或者吐槽保存
const (
	address     = "localhost:50051"
	defaultName = "world"
)

type SSSS struct {
	Name string
}

type SSSSbak struct {
	Data interface{}
}

// 问答或则吐槽
func WenDaOrTuCao(strnickName, stravatarUrl, strdata string, w http.ResponseWriter) {
	glog.Info("strnickName, stravatarUrl, strparam", strnickName, stravatarUrl, strdata)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	name := strnickName + "☢" + stravatarUrl + "☢" + strdata
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

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
	bbw := PswEncrypt(string(b))
	// bbw := PswEncrypt("abcsfsssfsfsfsfsfsfsfsfs3535353535353533")
	// base64
	encodeString := base64.StdEncoding.EncodeToString([]byte(bbw))
	ddd := SSSSbak{
		Data: encodeString,
	}
	bb, _ := json.Marshal(ddd)
	log.Printf("Greeting b: %s", string(bb))
	//	fmt.Fprint(w, string(bb))
	fmt.Fprint(w, string(bb))
	return
}
