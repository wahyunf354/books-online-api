package books

import (
	//"books_online_api/business/book_details"
	//"books_online_api/business/image_books"
	//"books_online_api/business/book_types"

	"context"
	"mime/multipart"
	"time"
)

type Domain struct {
	Keyword string

	Id int
	Title string
	UserId int
	Description string
	UrlBook string
	FileBook *multipart.FileHeader
	Price int
	BookTypeId int
	PageCount int
	BookDetailId int
	UrlCover []string
	FileCover []*multipart.FileHeader

	//BookType *book_types.Domain
	//BookDetail *book_details.Domain
	//ImageBooks []*image_books.Domain

	BookType interface{}
	BookDetail interface{}
	ImageBooks interface{}

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt interface{}
}

type Usecase interface {
	CreateBook(ctx context.Context, domain Domain) (Domain, error)
	GetBooks(ctx context.Context, domain Domain ) ([]Domain, error)
	GetOneBook(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	CreateBook(ctx context.Context, domain Domain) (Domain, error)
	GetBooks(ctx context.Context, domain Domain) ([]Domain, error)
	GetOneBook(ctx context.Context, domain Domain) (Domain, error)
}

type Localy interface {
	CreateBook(ctx context.Context, books Domain) (Domain, error)
}

type DetailRepository interface {
	CreateBook(ctx context.Context, books Domain) (Domain, error)
}

type ImageBooksLocaly interface {
	UploadImages(ctx context.Context, books Domain) (Domain, error)
}

type ImageBooksRepository interface {
	UploadImages(ctx context.Context, books Domain) (Domain, error)
}
