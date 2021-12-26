package entities

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

type OrderUseCase interface {
	GetAll(ctx context.Context) ([]OrderDomain, error)
	GetByCustomerId(ctx context.Context, customer_id int) ([]OrderDomain, error)
	GetById(ctx context.Context, id int) (OrderDomain, error)
}

type OrderRepository interface {
	GetAll(ctx context.Context) (res []OrderDomain, err error)
	GetByCustomerId(ctx context.Context, id int) ([]OrderDomain, error)
	GetById(ctx context.Context, id int) (OrderDomain, error)
}
