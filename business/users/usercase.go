package users

import (
	"context"
	"errors"
	"time"
)

type UserUsecase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUsecase(repo Repository, timeOut time.Duration) Usecase {
	return &UserUsecase{
		Repo:           repo,
		contextTimeout: timeOut,
	}
}

// core business login
func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("email empty")
	}

	if domain.Password == "" {
		return Domain{}, errors.New("password empty")
	}

	// var err error
	// domain.Password, err = encrypt.Hash(domain.Password)
	// if err != nil {
	// 	return Domain{}, err
	// }

	user, err := uc.Repo.Login(ctx, domain.Email, domain.Password)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUsecase) Register(ctx context.Context, userDomain *Domain) error {
	err := uc.Repo.Register(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}
