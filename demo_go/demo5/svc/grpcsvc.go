package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"learn-kit/demo5/proto"
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

func NewGrpcServer(endpoint EndpointSet, middlewares ...endpoint.Middleware) proto.MyServiceServer {
	for i := range middlewares {
		endpoint.echo = middlewares[i](endpoint.echo)
	}

	return &grpcServer{
		grpctransport.NewServer(
			endpoint.echo,
			decodeEchoRequest,
			encodeEchoResponse,
			//options...,
		),
	}
}

func decodeEchoRequest(ctx context.Context, i interface{}) (request interface{}, err error) {
	req := i.(*proto.EchoRequest)
	return EchoRequest{req.Ping}, nil
}

func encodeEchoResponse(ctx context.Context, i interface{}) (response interface{}, err error) {
	resp := i.(EchoResponse)
	return &proto.EchoResponse{Pong: resp.Pong}, nil
}
