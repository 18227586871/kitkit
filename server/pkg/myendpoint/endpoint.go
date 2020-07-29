package myendpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"micro_services/server/proto"
)

// 业务逻辑
type MyService struct{}

func (svc MyService) Echo(ctx context.Context, req *proto.EchoRequest) (resp *proto.EchoResponse, err error) {
	return &proto.EchoResponse{Pong: "Echo: " + req.Ping}, nil
}

func MakeEchoEndpoint(svc proto.MyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*proto.EchoRequest)
		return svc.Echo(ctx, req)
	}
}

type EndpointSet struct {
	Echo endpoint.Endpoint
}
