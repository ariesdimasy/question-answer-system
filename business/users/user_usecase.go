package users

import (
	"context"
	"time"
)

type UserUsecase struct {
	contextTimeout time.Duration
	repository     UserRepository
}

func NewUserUsecase(repo UserRepository, timeout time.Duration) UserUseCaseDomain {
	return &UserUsecase{
		contextTimeout: timeout,
		repository:     repo,
	}
}

func (uc *UserUsecase) GetAllUser(ctx context.Context) ([]UserDomain, error) {
	// timeout
	users, err := uc.repository.GetAllUser(ctx)
	if err != nil {
		return []UserDomain{}, err
	}
	return users, nil
}

func (uc *UserUsecase) GetUserById(ctx context.Context, id string) (UserDomain, error) {
	// timeout
	user, err := uc.repository.GetUserById(ctx, id)
	if err != nil {
		return UserDomain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (UserDomain, error) {
	user, err := uc.repository.Login(ctx, email, password)
	if err != nil {
		return UserDomain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) Register(ctx context.Context, userReg UserRegisterDomain) (UserDomain, error) {
	user, err := uc.repository.Register(ctx, userReg)
	if err != nil {
		return UserDomain{}, err
	}
	return user, nil
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id string) (UserDomain, error) {
	user, err := uc.repository.DeleteUser(ctx, id)
	if err != nil {
		return UserDomain{}, err
	}
	return user, nil
}
