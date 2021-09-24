package users

import (
	"context"
	"time"
)

type Domain struct {
	Id           int
	FirstName    string
	LastName     string
	Email        string
	Role         int8 // 1 for reader, 2 for writer
	Birth        time.Time
	Gender       string
	Token        string
	Password     string
	ConfPassword string
	HashPassword string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Usecase interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, email string, password string) (Domain, error)
}
