package main

import (
	"flag"
	"micro_services/server/pkg/mytransport"
	"net"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"

	"micro_services/server/pkg/myendpoint"
	"micro_services/server/proto"
)

func main() {
	var (
		grpcAddr            = flag.String("grpc-addr", ":8081", "gRPC listen address")
		logger              = log.NewLogfmtLogger(os.Stderr)
		loggingMiddleware   = myendpoint.LoggingMiddleware(logger)
		limiter             = rate.NewLimiter(rate.Every(time.Second), 100)
		rateLimitMiddleware = myendpoint.RateLimitMiddleware(limiter)
	)
	flag.Parse()

	service := myendpoint.MyService{}
	endpointSet := myendpoint.EndpointSet{Echo: myendpoint.MakeEchoEndpoint(service)}

	grpcServer := mytransport.NewGrpcServer(endpointSet, loggingMiddleware, rateLimitMiddleware)
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	proto.RegisterMyServiceServer(baseServer, grpcServer)

	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log(err)
	}

	logger.Log(baseServer.Serve(grpcListener))
}
