package service

import (
	"github.com/JloMkAaA/SiteForPractic"
	"github.com/JloMkAaA/SiteForPractic/pkg/repository"
)

type Authorization interface {
	CreateUser(user SiteForPractic.User) (int, error)
	GenerateToken(phone_number int, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UserControl interface {
	GetUserById(userIdAuth, userId int) (SiteForPractic.User, error)
	Delete(userIdAuth, userId int) error
	Update(userIdAuth, userId int, input SiteForPractic.UpdateUser) error
}

type Service struct {
	Authorization
	UserControl
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UserControl:   NewUserService(repos.UserManipulated),
	}
}
