package entities

import (
	"context"
	"time"
)

type OrderDetail struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	OrderId    string    `json:"order_id"`
	TotalPrice string    `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderDetailUseCase interface {
	GetAll(ctx context.Context) ([]OrderDetail, error)
	GetByOrderId(ctx context.Context, customer_id int) ([]OrderDetail, error)
	GetByOrderDetailId(ctx context.Context, id int) (OrderDetail, error)
}

type OrderDetailRepository interface {
	GetAll(ctx context.Context) (res []OrderDetail, err error)
	GetByOrderId(ctx context.Context, id int) ([]OrderDetail, error)
	GetByOrderDetailId(ctx context.Context, id int) (OrderDetail, error)
}
