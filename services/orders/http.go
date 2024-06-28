package main

import (
	"log"
	"net/http"

	handler "github.com/neerrrajj/oms/services/orders/handler/orders"
	"github.com/neerrrajj/oms/services/orders/service"
)

type hTTPServer struct {
	addr string
}

func NewHttpServer(addr string) *hTTPServer {
	return &hTTPServer{
		addr: addr,
	}
}

func (s *hTTPServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handler.NewOrdersHttpHandler(orderService)
	orderHandler.RegisterRouter(router)

	log.Println("starting http server on:", s.addr)
	return http.ListenAndServe(s.addr, router)
}
