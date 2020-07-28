package main

type MyService interface {
	Echo(req EchoRequest) EchoResponse
}

type myService struct{}

func (svc myService) Echo(req EchoRequest) EchoResponse {
	return EchoResponse{"Echo: " + req.Ping}
}

