package google_books_test

import (
	"books_online_api/business/google_books"
	_mocksGoogleBooks "books_online_api/business/google_books/mocks"
	"books_online_api/controllers"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var googleBooksThirdParts _mocksGoogleBooks.ThirdPartyGoogleBooks

var googleBooksService google_books.Usecase
var googleBooksDomains []google_books.Domain

func setup() {
	googleBooksService = google_books.NewGoogleBookThirtPartUsecase(&googleBooksThirdParts, time.Hour*1)
	domain := google_books.Domain{}
	googleBooksDomains = []google_books.Domain{domain, domain}
}

func TestGoogleBookUserCase_SearchBooks(t *testing.T) {
	t.Run("Test case 1 | empty keyword", func(t *testing.T) {
		setup()
		googleBooksThirdParts.On("SearchBooks", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(googleBooksDomains, nil)
		_, err := googleBooksService.SearchBooks(context.Background(), "", 10, 10)
		assert.NotNil(t, err)
		assert.Equal(t, controllers.KEYWORD_EMPTY, err)
	})

	t.Run("Test case 2 | success get books", func(t *testing.T) {
		setup()
		googleBooksThirdParts.On("SearchBooks", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(googleBooksDomains, nil)
		_, err := googleBooksService.SearchBooks(context.Background(), "Golang", 10, 10)
		assert.Nil(t, err)
	})

	t.Run("Test case 2 | max result 0", func(t *testing.T) {
		setup()
		googleBooksThirdParts.On("SearchBooks", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(googleBooksDomains, nil)
		_, err := googleBooksService.SearchBooks(context.Background(), "Golang", 10, 0)
		assert.Nil(t, err)
	})
}
