package grpcsvc

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	endpoint2 "micro_services/server/endpoint"
	proto "micro_services/server/pb"
)

// grpc服务
type grpcServer struct {
	echo grpctransport.Handler
}

func (s grpcServer) Echo(ctx context.Context, request *proto.EchoRequest) (*proto.EchoResponse, error) {
	_, resp, err := s.echo.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*proto.EchoResponse), nil
}

// 应用middleware
func NewGrpcServer(endpoint endpoint2.EndpointSet, middlewares ...endpoint.Middleware) proto.MyServiceServer {
	for i := range middlewares {
		endpoint.Echo = middlewares[i](endpoint.Echo)
	}

	return &grpcServer{
		grpctransport.NewServer(
			endpoint.Echo,
			decodeEchoRequest,
			encodeEchoResponse,
		),
	}
}

func decodeEchoRequest(ctx context.Context, i interface{}) (request interface{}, err error) {
	return i.(*proto.EchoRequest), nil
}

func encodeEchoResponse(ctx context.Context, i interface{}) (response interface{}, err error) {
	return i.(*proto.EchoResponse), nil
}
