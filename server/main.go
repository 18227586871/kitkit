package main

import (
	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"micro_services/server/endpoint"
	"micro_services/server/grpcsvc"
	"micro_services/server/middleware"
	proto "micro_services/server/pb"
	"net"
	"os"
	"time"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	loggingMiddleware := middleware.LoggingMiddleware(logger)
	limiter := rate.NewLimiter(rate.Every(time.Second), 100)
	rateLimitMiddleware := middleware.RateLimitMiddleware(limiter)

	service := endpoint.MyService{}
	endpointSet := endpoint.EndpointSet{Echo: endpoint.MakeEchoEndpoint(service)}

	grpcServer := grpcsvc.NewGrpcServer(endpointSet, loggingMiddleware, rateLimitMiddleware)
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	proto.RegisterMyServiceServer(baseServer, grpcServer)

	grpcListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Log(err)
	}

	logger.Log(baseServer.Serve(grpcListener))
}
