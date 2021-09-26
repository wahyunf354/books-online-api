package book_types

import (
	"context"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"time"
)

type BookTypeDomain struct {
	Id int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name string
	Unit string
	DeletedAt gorm.DeletedAt
	Role int8
}

type Category struct {
	Id int
	CreatedAt time.Time
	UpdatedAt time.Time
	Title string
}

type ImageBook struct {
	Id int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Url string
	BookId int
}

type Books struct {
	Id int
	Title      string
	BookTypeId int
	CategoryId int
	Price      int
	UserId     int
	BookType   *BookTypeDomain
	Category   *Category
	Images     []*ImageBook

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	CreateBookType(ctx context.Context, bookType BookTypeDomain) (BookTypeDomain, error)
}

type Repository interface {
	CreateBookType(ctx context.Context, bookType BookTypeDomain) (BookTypeDomain, error)
}

type Controller interface {
	CreateBookType(c echo.Context) error
}
