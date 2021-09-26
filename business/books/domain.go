package books

import (
	"context"
	"mime/multipart"
	"time"
)

type Domain struct {
	Id int
	Title string
	UserId int
	Description string
	UrlBook string
	Price int
	BookTypeId int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt time.Time
	FileBook multipart.FileHeader
}

type Usecase interface {
	CreateBook(ctx context.Context, books Domain) (Domain, error)
}

type Repository interface {
	CreateBook(ctx context.Context, books Domain) (Domain, error)
}

type Localy interface {
	CreateBook(ctx context.Context, books Domain) (Domain, error)
}