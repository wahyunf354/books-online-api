package book_types

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

type Domain struct {
	Id int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string
	Unit string
	DeletedAt gorm.DeletedAt
	Role int8
}

type Usecase interface {
	CreateBookType(ctx context.Context, bookType Domain) (Domain, error)
}

type Repository interface {
	CreateBookType(ctx context.Context, bookType Domain) (Domain, error)
}

type Controller interface {
	CreateBookType(c echo.Context) error
}
