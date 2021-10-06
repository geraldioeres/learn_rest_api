package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	Email     string
	Address   string
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	Register(ctx context.Context, data *Domain) error
}

type Repository interface{
	Login(ctx context.Context, email string, password string) (Domain, error)
	Register(ctx context.Context, data *Domain) error
}