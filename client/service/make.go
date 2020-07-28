package client

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakePingEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return PingReq{
			Ping: RpcResp(),
		}, nil
	}
}

// 解析请求过来的参数 request
func decodePingReq(ctx context.Context, r *http.Request) (resp interface{}, err error) {
	var request PingReq
	return request, nil
}

// 写回页面通用方法
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint) http.Handler {
	r := mux.NewRouter()
	// 路由
	r.Methods("GET").Path("/ping").Handler(httptransport.NewServer(
		endpoint,
		decodePingReq,
		encodeResponse,
	))

	return r
}
