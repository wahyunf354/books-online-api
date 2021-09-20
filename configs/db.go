package configs

import (
	"books_online_api/models/books"
	"books_online_api/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "Book$online123:Book$online123@tcp(127.0.0.1:3306)/books_online?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Connection DB Faild")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(
		&users.User{},
		&books.Book{},
		&books.BookType{},
		&books.Category{},
		&books.ImageBooks{},
		&books.BookDetails{},
	)

	DB.Migrator().AlterColumn(&users.User{}, "Email")
}
