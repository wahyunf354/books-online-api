package books_test

import (
	"books_online_api/business/books"
	_mocksBooks "books_online_api/business/books/mocks"
	"books_online_api/controllers"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mime/multipart"
	"testing"
	"time"
)

var booksRepository _mocksBooks.Repository
var booksLocaly _mocksBooks.Localy

var booksService books.Usecase
var bookDomain books.Domain
var bookDomains []books.Domain

func setupCreateBooks() {
	booksService = books.NewBookUsecase(&booksLocaly, &booksRepository, time.Hour*1)
	bookDomain = books.Domain{
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
		setupCreateBooks()
		booksLocaly.On("CreateBook", mock.Anything, mock.Anything).Return(bookDomain, nil).Once()

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
		setupCreateBooks()
		booksLocaly.On("CreateBook", mock.Anything, mock.Anything).Return(bookDomain, nil).Once()

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
		assert.Equal(t, controllers.EMPTY_TITLE, err)
	})

	t.Run("Test case 3 | Description Empty", func(t *testing.T) {
		setupCreateBooks()
		booksLocaly.On("CreateBook", mock.Anything, mock.Anything).Return(bookDomain, nil).Once()

		book, err := booksService.CreateBook(context.Background(), books.Domain{
			Title:        "Golang",
			UserId:       13,
			Description:  "",
			FileBook:     nil,
			Price:        25000,
			BookTypeId:   1,
			PageCount:    200,
			FileCover:    []*multipart.FileHeader{},
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", book.Description)
		assert.Equal(t, controllers.EMPTY_DESCRIPTION, err)
	})

}

func setupGetOneBook() {
	booksService = books.NewBookUsecase(&booksLocaly, &booksRepository, time.Hour*1)
	bookDomain = books.Domain{
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

func TestBookUsecase_GetOneBook(t *testing.T) {
	t.Run("Test Case 1 | valid get one book", func(t *testing.T) {
		setupGetOneBook()
		booksRepository.On("GetOneBook", mock.Anything, mock.Anything).Return(bookDomain, nil)
		book, err := booksService.GetOneBook(context.Background(), books.Domain{Id: 2})
		assert.Nil(t, err)
		assert.NotNil(t, book)
		assert.Equal(t, "Golang", book.Title)
	})
	t.Run("Test Case 2 | Id empty", func(t *testing.T) {
		setupGetOneBook()
		booksRepository.On("GetOneBook", mock.Anything, mock.Anything).Return(bookDomain, nil)
		_, err := booksService.GetOneBook(context.Background(), books.Domain{Id:0})
		assert.NotNil(t, err)
		assert.Equal(t, controllers.ID_EMPTY, err)
	})
}