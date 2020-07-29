package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"micro_service/client/pkg/model"
	"micro_service/client/pkg/myendpoint"
)

var address string = "localhost:9000"

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var svc model.Service
	svc = model.ServiceStruct{}
	ep := myendpoint.MakePingEndpoint(svc)

	r := myendpoint.MakeHttpHandler(ctx, ep)
	go func() {
		log.Println("Http Server start at port", address)
		handler := r
		errChan <- http.ListenAndServe(address, handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	log.Println(<-errChan)
}
