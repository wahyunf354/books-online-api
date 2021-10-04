package book_types_test

import (
	"books_online_api/business/book_types"
	_mocksBookTypesRepository "books_online_api/business/book_types/mocks"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"testing"
	"time"
)

var bookTypeRepository _mocksBookTypesRepository.Repository

var bookTypeService book_types.Usecase
var bookTypeDomain book_types.Domain

func setup() {
	bookTypeService = book_types.NewBooksUseCase(&bookTypeRepository, time.Hour*1)
	bookTypeDomain = book_types.Domain{
		Id:        1,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		Name:      "Gratis",
		Unit:      "Selamanya",
		DeletedAt: gorm.DeletedAt{},
	}
}

func TestBookTypesUseCase_CreateBookType(t *testing.T) {
	t.Run("Test case 1 | success create book type", func(t *testing.T) {
		setup()
		bookTypeRepository.On("CreateBookType", mock.Anything, mock.Anything).Return(bookTypeDomain, nil)

		bookType, err := bookTypeService.CreateBookType(context.Background(), book_types.Domain{
			Name:      "Gratis",
			Unit:      "Selamanya",
		})
		assert.Nil(t, err)
		assert.Equal(t, "Gratis", bookType.Name)
		assert.Equal(t, "Selamanya", bookType.Unit)
	})

	t.Run("Test case 2 | Name empty", func(t *testing.T) {
		setup()
		bookTypeRepository.On("CreateBookType", mock.Anything, mock.Anything).Return(bookTypeDomain, nil)

		bookType, err := bookTypeService.CreateBookType(context.Background(), book_types.Domain{
			Name:      "",
			Unit:      "Selamanya",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", bookType.Name)
		assert.Equal(t, errors.New("name book type empty"), err)
	})

	t.Run("Test case 3 | Unit Empty", func(t *testing.T) {
		setup()
		bookTypeRepository.On("CreateBookType", mock.Anything, mock.Anything).Return(bookTypeDomain, nil)

		bookType, err := bookTypeService.CreateBookType(context.Background(), book_types.Domain{
			Name:      "Gratis",
			Unit:      "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, "", bookType.Unit)
		assert.Equal(t, errors.New("unit book type empty"), err)
	})
}