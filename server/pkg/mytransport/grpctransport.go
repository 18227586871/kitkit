package mytransport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	"micro_services/server/pkg/myendpoint"
	"micro_services/server/proto"
)

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

func NewGrpcServer(endpointSet myendpoint.EndpointSet, middlewares ...endpoint.Middleware) proto.MyServiceServer {
	for i := range middlewares {
		endpointSet.Echo = middlewares[i](endpointSet.Echo)
	}

	return &grpcServer{
		grpctransport.NewServer(
			endpointSet.Echo,
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
