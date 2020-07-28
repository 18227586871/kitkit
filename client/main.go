package main

import (
	"context"
	"fmt"
	client "micro_service/client/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var svc client.Service
	svc = client.ServiceStruct{}
	endpoint := client.MakePingEndpoint(svc)

	r := client.MakeHttpHandler(ctx, endpoint)
	go func() {
		fmt.Println("Http Server start at port:9000")
		handler := r
		errChan <- http.ListenAndServe(":9000", handler)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
