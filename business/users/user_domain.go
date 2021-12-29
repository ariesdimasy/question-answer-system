package users

import (
	"context"
	"time"
)

type UserDomain struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRegisterDomain struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginDomain struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUseCaseDomain interface {
	GetAllUser(ctx context.Context) ([]UserDomain, error)
	GetUserById(ctx context.Context, id string) (UserDomain, error)
	Login(ctx context.Context, userLogin UserLoginDomain) (UserDomain, error)
	Register(ctx context.Context, user UserRegisterDomain) (UserDomain, error)
	DeleteUser(ctx context.Context, id string) (UserDomain, error)
}

type UserRepository interface {
	GetAllUser(ctx context.Context) (res []UserDomain, err error)
	GetUserById(ctx context.Context, id string) (res UserDomain, err error)
	Login(ctx context.Context, userLogin UserLoginDomain) (res UserDomain, err error)
	Register(ctx context.Context, user UserRegisterDomain) (res UserDomain, err error)
	DeleteUser(ctx context.Context, id string) (res UserDomain, err error)
}
