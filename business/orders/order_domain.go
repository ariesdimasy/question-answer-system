package orders

import (
	"context"
	"time"
)

type OrderDomain struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	CustomerId string    `json:"customer_id"`
	TotalPrice string    `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateOrderDomain struct {
	CustomerId string `json:"customer_id"`
	TotalPrice string `json:"total_price"`
}

type UpdateOrderDomain struct {
	Id         uint   `json:"id" gorm:"primary_key"`
	CustomerId string `json:"customer_id"`
	TotalPrice string `json:"total_price"`
}

type OrderUseCaseDomain interface {
	GetAllOrders(ctx context.Context) ([]OrderDomain, error)
	GetOrderByCustomerId(ctx context.Context, customer_id int) ([]OrderDomain, error)
	GetOrderById(ctx context.Context, id int) (OrderDomain, error)
	CreateOrder(ctx context.Context, createOrder CreateOrderDomain) (OrderDomain, error)
	UpdateOrder(ctx context.Context, updateOrder UpdateOrderDomain) (OrderDomain, error)
	DeleteOrder(ctx context.Context, id int) (OrderDomain, error)
}

type OrderRepository interface {
	GetAllOrders(ctx context.Context) (res []OrderDomain, err error)
	GetOrderByCustomerId(ctx context.Context, customer_id int) (res []OrderDomain, err error)
	GetOrderById(ctx context.Context, id int) (res OrderDomain, err error)
	CreateOrder(ctx context.Context, createOrder CreateOrderDomain) (res OrderDomain, err error)
	UpdateOrder(ctx context.Context, updateOrder UpdateOrderDomain) (res OrderDomain, err error)
	DeleteOrder(ctx context.Context, id int) (res OrderDomain, err error)
}
