package demo

import (
	"context"
	proto "demo_go/pb"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

func MakeAdd(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return PingReq{
			Ping: RpcResp(),
		}, nil
	}
}

func decodeReq(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	// 逻辑
	return PingResp{
		Pong: "testResp",
	}, nil
}

func encodeArithmeticResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/ping").Handler(httptransport.NewServer(
		endpoint,
		decodeReq,
		encodeArithmeticResponse,
	))

	return r

}

func RpcResp() string {
	conn, err := grpc.Dial("10.221.113.184:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewMyServiceClient(conn)

	if len(os.Args) > 1 {
		fmt.Println(os.Args[1])
	}
	r, err := c.Echo(context.Background(), &proto.EchoRequest{Ping: "ping"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Pong + "687989748"
}

//func NewTokenBucketLimitterWithBuildIn(bkt *rate.Limiter) endpoint.Middleware {
//	return func(next endpoint.Endpoint) endpoint.Endpoint {
//		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
//			if !bkt.Allow() {
//				return nil, errors.New("limit")
//			}
//			return next(ctx, request)
//		}
//	}
//}
