package main

import (
	"context"
	"demo_go/demo"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var svc demo.Service
	svc = demo.ServiceStruct{}
	endpoint := demo.MakeAdd(svc)

	//limiter := rate.NewLimiter(rate.Every(time.Second), 3)
	//
	//endpoint = demo.NewTokenBucketLimitterWithBuildIn(limiter)(endpoint)

	r := demo.MakeHttpHandler(ctx, endpoint)
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
