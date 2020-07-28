package main
//
//import (
//	proto "demo_go/pb"
//	"fmt"
//	"golang.org/x/net/context"
//	"google.golang.org/grpc"
//	"log"
//	"os"
//)
//
//const (
//	address = "10.221.113.184:8081"
//)
//
//func main() {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := proto.NewMyServiceClient(conn)
//
//	if len(os.Args) > 1 {
//		fmt.Println(os.Args[1])
//	}
//	r, err := c.Echo(context.Background(), &proto.EchoRequest{Ping: "ping"})
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	log.Printf("####### get server Greeting response: %s", r.Pong)
//}
//
//func RpcResp() string {
//	conn, err := grpc.Dial(address, grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("did not connect: %v", err)
//	}
//	defer conn.Close()
//	c := proto.NewMyServiceClient(conn)
//
//	if len(os.Args) > 1 {
//		fmt.Println(os.Args[1])
//	}
//	r, err := c.Echo(context.Background(), &proto.EchoRequest{Ping: "ping"})
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	return r.Pong
//}
