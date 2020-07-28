package main

import (
	"demo_go/demo5/proto"
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	runGrpcServer()

}

func runGrpcServer() {
	logger := log.NewLogfmtLogger(os.Stderr)
	loggingMiddleware := LoggingMiddleware(logger)
	//limiter := rate.NewLimiter(rate.Every(time.Second), 1)
	//rateLimitMiddleware := RateLimitMiddleware(limiter)

	var service MyService
	service = myService{}
	endpointSet := EndpointSet{echo: MakeEchoEndpoint(service)}

	grpcServer := NewGrpcServer(endpointSet, loggingMiddleware)
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	proto.RegisterMyServiceServer(baseServer, grpcServer)

	grpcListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Log(err)
	}

	logger.Log(baseServer.Serve(grpcListener))
}
