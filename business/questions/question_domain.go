package questions

import (
	"context"
	"time"
)

type QuestionDomain struct {
	Id        uint      `json:"id" gorm:"primary_key"`
	Question  string    `json:"question"`
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type QuestionAddDomain struct {
	Question string `json:"question"`
	TestId   int    `json:"test_id"`
	Value    int    `json:"value"`
}

type QuestionUpdateDomain struct {
	Id       uint   `json:"id" gorm:"primary_key"`
	Question string `json:"question"`
	TestId   int    `json:"test_id"`
	Value    int    `json:"value"`
}

type QuestionUseCaseDomain interface {
	GetAllQuestionbyTestId(ctx context.Context, test_id int) ([]QuestionDomain, error)
	GetQuestionById(ctx context.Context, id string) (QuestionDomain, error)
}

type QuestionRepository interface {
	GetAllQuestionbyTestId(ctx context.Context, test_id int) (res []QuestionDomain, err error)
	GetQuestionById(ctx context.Context, id string) (res QuestionDomain, err error)
}
