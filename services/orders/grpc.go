package main

import (
	"log"
	"net"

	handler "github.com/neerrrajj/oms/services/orders/handler/orders"
	"github.com/neerrrajj/oms/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	//registering our grpc services
	orderService := service.NewOrderService()
	handlerService := handler.NewOrdersGrpcHandler(orderService)
	handler.NewGrpcOrdersService(grpcServer, handlerService)

	log.Println("starting our gRPC server on", s.addr)
	return grpcServer.Serve(lis)
}
