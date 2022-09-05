package questions

import (
	"context"
	"time"
)

type QuestionUsecase struct {
	contextTimeout time.Duration
	repository     QuestionRepository
}

func NewQuestionUsecase(repo QuestionRepository, timeout time.Duration) QuestionUseCaseDomain {
	return &QuestionUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (qc *QuestionUsecase) GetAllQuestionbyTestId(ctx context.Context, test_id int) ([]QuestionDomain, error) {
	// timeout
	questions, err := qc.repository.GetAllQuestionbyTestId(ctx, test_id)
	if err != nil {
		return []QuestionDomain{}, err
	}
	return questions, nil
}

func (qc *QuestionUsecase) GetQuestionById(ctx context.Context, id string) (QuestionDomain, error) {
	// timeout
	question, err := qc.repository.GetQuestionById(ctx, id)
	if err != nil {
		return QuestionDomain{}, err
	}
	return question, nil
}
