package books_test

import (
	"books_online_api/business/books"
	_mocksBooks "books_online_api/business/books/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"testing"
	"time"
)

var booksRepository _mocksBooks.Repository
var booksLocaly _mocksBooks.Localy

var booksService books.Usecase
var booksDomain books.Domain

func setupCreateBook() {
	booksService = books.NewBookUsecase(&booksLocaly, &booksRepository, time.Hour*1)
	booksDomain = books.Domain{
		Keyword:      "",
		Id:           1,
		Title:        "Golang",
		UserId:       2,
		Description:  "Golang belajar",
		UrlBook:      "http://file.jpg",
		FileBook:     nil,
		Price:        0,
		BookTypeId:   0,
		PageCount:    0,
		BookDetailId: 0,
		UrlCover:     nil,
		FileCover:    nil,
		BookType:     nil,
		BookDetail:   nil,
		ImageBooks:   nil,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
	}
}

func TestBookUsecase_CreateBook(t *testing.T) {
	t.Run("Test case 1 | Create Books Valid", func(t *testing.T) {
		setupCreateBook()
		booksLocaly.On("CreateBook", mock.Anything, mock.Anything).Return(booksDomain, nil).Once()

		book, err := booksService.CreateBook(context.Background(), books.Domain{
			Title:        "Golang",
			UserId:       13,
			Description:  "Membahas tentang algoritma",
			FileBook:     nil,
			Price:        25000,
			BookTypeId:   1,
			PageCount:    200,
			FileCover:    []*multipart.FileHeader{},
		})
		assert.Nil(t, err)
		assert.Equal(t, "Golang", book.Title)
	})

	t.Run("Test case 2 | Title Empty", func(t *testing.T) {
		setupCreateBook()
		booksLocaly.On("CreateBook", mock.Anything, mock.Anything).Return(booksDomain, nil).Once()

		book, err := booksService.CreateBook(context.Background(), books.Domain{
			Title:        "",
			UserId:       13,
			Description:  "Membahas tentang algoritma",
			FileBook:     nil,
			Price:        25000,
			BookTypeId:   1,
			PageCount:    200,
			FileCover:    []*multipart.FileHeader{},
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", book.Title)
		assert.Equal(t, errors.New("Title empty"), err)
	})

}