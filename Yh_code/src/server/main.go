//package main

//import (
//	"log"
//	"net"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	pb "github.com/freewebsys/grpc-go-demo/src/helloworld"
//	"google.golang.org/grpc/reflection"
//	"fmt"
//)

//const (
//	port = ":50051"
//)

//// server is used to implement helloworld.GreeterServer.
//type server struct{}

//// SayHello implements helloworld.GreeterServer
//func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//	fmt.Println("######### get client request name :"+in.Name)
//	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
//}

//func main() {
//	lis, err := net.Listen("tcp", port)
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	pb.RegisterGreeterServer(s, &server{})
//	// Register reflection service on gRPC server.
//	reflection.Register(s)
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}

//package main

//import (
//	"fmt"
//	"io"
//	"log"
//	"net"

//	"net/http"

//	consulapi "github.com/hashicorp/consul/api"
//)

//const RECV_BUF_LEN = 1024

//func consulCheck(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "consulCheck")
//}

//func registerServer() {

//	config := consulapi.DefaultConfig()
//	client, err := consulapi.NewClient(config)

//	if err != nil {
//		log.Fatal("consul client error : ", err)
//	}

//	checkPort := 8080

//	registration := new(consulapi.AgentServiceRegistration)
//	registration.ID = "serverNode_1"
//	registration.Name = "serverNode"
//	registration.Port = 9527
//	registration.Tags = []string{"serverNode"}
//	registration.Address = "127.0.0.1"
//	registration.Check = &consulapi.AgentServiceCheck{
//		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, checkPort, "/check"),
//		Timeout:                        "3s",
//		Interval:                       "5s",
//		DeregisterCriticalServiceAfter: "30s", //check失败后30秒删除本服务
//	}
//  注册服务操作
//	err = client.Agent().ServiceRegister(registration)

//	if err != nil {
//		log.Fatal("register server error : ", err)
//	}

//	http.HandleFunc("/check", consulCheck)
//	http.ListenAndServe(fmt.Sprintf(":%d", checkPort), nil)

//}

//func main() {
// 注册操作
//	go registerServer()

//	ln, err := net.Listen("tcp", "0.0.0.0:9527")

//	if nil != err {
//		panic("Error: " + err.Error())
//	}

//	for {
//		conn, err := ln.Accept()

//		if err != nil {
//			panic("Error: " + err.Error())
//		}

//		go EchoServer(conn)
//	}

//}

//func EchoServer(conn net.Conn) {
//	buf := make([]byte, RECV_BUF_LEN)
//	defer conn.Close()

//	for {
//		n, err := conn.Read(buf)
//		switch err {
//		case nil:
//			log.Println("get and echo:", "EchoServer "+string(buf[0:n]))
//			conn.Write(append([]byte("EchoServer "), buf[0:n]...))
//		case io.EOF:
//			log.Printf("Warning: End of data: %s\n", err)
//			return
//		default:
//			log.Printf("Error: Reading data: %s\n", err)
//			return
//		}
//	}
//}

package main

import (
	"fmt"
	"log"
	"net"
	"time"

	consulapi "github.com/hashicorp/consul/api"
)

const RECV_BUF_LEN = 1024

func main() {
	// 申请客户端操作
	client, err := consulapi.NewClient(consulapi.DefaultConfig())

	if err != nil {
		log.Fatal("consul client error : ", err)
	}

	// 数据处理操作 应用操作
	for {

		time.Sleep(time.Second * 3)
		var services map[string]*consulapi.AgentService
		var err error

		services, err = client.Agent().Services()

		if nil != err {
			log.Println("in consual list Services:", err)
			continue
		}

		if _, found := services["serverNode_1"]; !found {
			log.Println("serverNode_1 not found")
			continue
		}

		sendData(services["serverNode_1"])

	}
}

// 获取服务操作
// 1 主要是操作的
// 2 查找操作
// 3 数据操作，应用操作
func sendData(service *consulapi.AgentService) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", service.Address, service.Port))

	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	buf := make([]byte, RECV_BUF_LEN)
	i := 0
	for {
		i++
		msg := fmt.Sprintf("Hello World, %03d", i)
		n, err := conn.Write([]byte(msg))
		if err != nil {
			println("Write Buffer Error:", err.Error())
			break
		}

		n, err = conn.Read(buf)
		if err != nil {
			println("Read Buffer Error:", err.Error())
			break
		}
		log.Println("get:", string(buf[0:n]))

		//等一秒钟
		time.Sleep(time.Second)
	}
}
