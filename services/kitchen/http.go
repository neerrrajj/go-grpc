package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/neerrrajj/oms/services/common/genproto/orders"
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

	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/favicon.ico" {
			http.NotFound(w, r)
			return
		}

		c := orders.NewOrderServiceClient(conn)

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
		defer cancel()

		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 11,
			ProductID:  11,
			Quantity:   11,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		res, err := c.GetOrder(ctx, &orders.GetOrderRequest{
			CustomerID: 31,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		t := template.Must(template.New("orders").Parse(ordersTemplate))

		err = t.Execute(w, res.GetOrders())
		if err != nil {
			log.Fatalf("template error: %v", err)
		}
	})

	log.Println("starting server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
