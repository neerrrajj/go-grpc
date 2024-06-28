package handler

import (
	"context"

	"github.com/neerrrajj/oms/services/common/genproto/orders"
	"github.com/neerrrajj/oms/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewOrdersGrpcHandler(ordersService types.OrderService) *OrdersGrpcHandler {
	return &OrdersGrpcHandler{
		ordersService: ordersService,
	}
}

func NewGrpcOrdersService(grpcServer *grpc.Server, gRPCHandler *OrdersGrpcHandler) {
	orders.RegisterOrderServiceServer(grpcServer, gRPCHandler)
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {

	order := &orders.Order{
		OrderID:    42,
		CustomerID: 24,
		ProductID:  12,
		Quantity:   21,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}

	return res, nil
}

func (h *OrdersGrpcHandler) GetOrder(ctx context.Context, req *orders.GetOrderRequest) (*orders.GetOrderResponse, error) {
	o := h.ordersService.GetOrders(ctx)
	res := &orders.GetOrderResponse{
		Orders: o,
	}
	return res, nil
}
