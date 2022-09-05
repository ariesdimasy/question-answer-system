package tests

import (
	"context"
	"time"
)

type TestUsecase struct {
	contextTimeout time.Duration
	repository     TestRepository
}

func NewQuestionUsecase(repo TestRepository, timeout time.Duration) TestUseCaseDomain {
	return &TestUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (qc *TestUsecase) GetAllTest(ctx context.Context) ([]TestDomain, error) {
	// timeout
	tests, err := qc.repository.GetAllTest(ctx)
	if err != nil {
		return []TestDomain{}, err
	}
	return tests, nil
}

func (qc *TestUsecase) GetTestById(ctx context.Context, id int) (TestDomain, error) {
	// timeout
	test, err := qc.repository.GetTestById(ctx, id)
	if err != nil {
		return TestDomain{}, err
	}
	return test, nil
}
