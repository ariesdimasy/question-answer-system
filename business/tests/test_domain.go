package tests

import (
	"context"
	"time"
)

type TestDomain struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TestAddDomain struct {
	Title string `json:"title"`
}

type TestUpdateDomain struct {
	Id    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
}

type TestUseCaseDomain interface {
	GetAllTest(ctx context.Context) ([]TestDomain, error)
	GetTestById(ctx context.Context, id int) (TestDomain, error)
}

type TestRepository interface {
	GetAllTest(ctx context.Context) (res []TestDomain, err error)
	GetTestById(ctx context.Context, id int) (res TestDomain, err error)
}
