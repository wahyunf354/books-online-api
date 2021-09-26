package book_types

import (
	"books_online_api/business/book_types"
	"context"
	"gorm.io/gorm"
)

type MysqlBookTypesRepository struct {
	Conn *gorm.DB
}

func NewMysqlBookTypesRepository(conn *gorm.DB) book_types.Repository {
	return &MysqlBookTypesRepository{
		Conn: conn,
	}
}

func (rep *MysqlBookTypesRepository) CreateBookType(ctx context.Context, bookType book_types.BookTypeDomain) (book_types.BookTypeDomain, error) {

	newBookType := FromDomain(bookType)

	result := rep.Conn.Create(&newBookType)

	if result.Error != nil {
		return book_types.BookTypeDomain{}, result.Error
	}

	return newBookType.ToDomain(), nil
}
