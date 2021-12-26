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
	GetAll(ctx context.Context) ([]Order, error)
	GetByCustomerId(ctx context.Context, customer_id int) ([]Order, error)
	GetById(ctx context.Context, id int) (Order, error)
}

type OrderDetailRepository interface {
	GetAll(ctx context.Context) (res []Order, err error)
	GetByCustomerId(ctx context.Context, id int) ([]Order, error)
	GetById(ctx context.Context, id int) (Order, error)
}
