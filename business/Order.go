package entities

import (
	"context"
	"time"
)

type Order struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	CustomerId string    `json:"customer_id"`
	TotalPrice string    `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderUseCase interface {
	GetAll(ctx context.Context) ([]Order, error)
	GetByCustomerId(ctx context.Context, customer_id int) ([]Order, error)
	GetById(ctx context.Context, id int) (Order, error)
}

type BookRepository interface {
	GetAll(ctx context.Context) (res []Order, err error)
	GetByCustomerId(ctx context.Context, id int) ([]Order, error)
	GetById(ctx context.Context, id int) (Order, error)
}
