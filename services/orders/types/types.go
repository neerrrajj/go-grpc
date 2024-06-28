package types

import (
	"context"

	"github.com/neerrrajj/oms/services/common/genproto/orders"
)

// why do we have to create an interface when server and client interfaces are already generated?

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context) []*orders.Order
}
