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
	FileBook multipart.FileHeader
	PageCount int
	BookDetailId int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt interface{}
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

type DetailRepository interface {
	CreateBook(ctx context.Context, books Domain) (Domain, error)
}