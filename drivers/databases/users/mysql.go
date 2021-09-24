package users

import (
	"books_online_api/business/users"
	"context"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	user := Users{}
	user.FirstName = domain.FirstName
	user.LastName = domain.LastName
	user.Email = domain.Email
	user.Gender = domain.Gender
	user.Role = domain.Role
	user.HashPassword = domain.HashPassword
	user.Birth = domain.Birth

	result := rep.Conn.Create(&user)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}

func (rep *MysqlUserRepository) Login(ctx context.Context, email string, password string) (users.Domain, error) {

	var user Users
	result := rep.Conn.First(&user, "email = ? AND password = ?", email, password)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return user.ToDomain(), nil
}
