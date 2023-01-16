package repository

import (
	"github.com/JloMkAaA/SiteForPractic"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user SiteForPractic.User) (int, error)
	GetUser(phone_number int, password string) (SiteForPractic.User, error)
}

type UserManipulated interface {
	GetUserById(userIdAuth, userId int) (SiteForPractic.User, error)
	Delete(userIdAuth, userId int) error
	Update(userIdAuth, userId int, input SiteForPractic.UpdateUser) error
}

type Repository struct {
	Authorization
	UserManipulated
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:   NewAuthPostgres(db),
		UserManipulated: NewUserPostgres(db),
	}
}
