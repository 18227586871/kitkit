package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type EchoRequest struct {
	Ping string `json:"ping"`
}

type EchoResponse struct {
	Pong string `json:"pong"`
}

func MakeEchoEndpoint(svc MyService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(EchoRequest)
		response = svc.Echo(req)
		return
	}
}

type EndpointSet struct {
	echo endpoint.Endpoint
}
