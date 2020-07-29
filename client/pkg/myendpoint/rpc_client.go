package myendpoint

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"
	proto "micro_service/client/pb"
)

var address string = "10.221.113.184:8081"

// 调用rpc
func RpcResp() string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewMyServiceClient(conn)

	if len(os.Args) > 1 {
		log.Println(os.Args[1])
	}
	r, err := c.Echo(context.Background(), &proto.EchoRequest{Ping: "ping"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Pong
}
