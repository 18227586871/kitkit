package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	"micro_service/api/pb"
	severendpoint "micro_service/internal/app/server/endpoint"
)

type grpcServer struct {
	echo grpctransport.Handler
}

func (s grpcServer) Echo(ctx context.Context, request *pb.EchoRequest) (*pb.EchoResponse, error) {
	_, resp, err := s.echo.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.EchoResponse), nil
}

// 统一处理endpoint和middleware并返回
func NewGrpcServer(endpointSet severendpoint.EndpointSet, middlewares ...endpoint.Middleware) pb.MyServiceServer {
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

func decodeEchoRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request.(*pb.EchoRequest), nil
}

func encodeEchoResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response.(*pb.EchoResponse), nil
}
