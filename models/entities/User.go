package entities

import (
	"context"
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUseCase interface {
	GetAll(ctx context.Context) ([]User, error)
	GetById(ctx context.Context, id string) (User, error)
}

type UserRepository interface {
	GetAll(ctx context.Context) (
		res []User, err error)
	GetById(ctx context.Context, id string) (User, error)
}
