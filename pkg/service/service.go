package service

import (
	"github.com/JloMkAaA/SiteForPractic"
	"github.com/JloMkAaA/SiteForPractic/pkg/repository"
)

type Authorization interface {
	CreateUser(user SiteForPractic.User) (int, error)
	GenerateToken(phone_number uint64, password string) (string, error)
	ParseToken(token string) (int, error)
	NewRefreshToken() (string, error)
}

type UserControl interface {
	GetUserById(userId int) (SiteForPractic.User, error)
	Delete(userId int) error
	Update(userId int, input SiteForPractic.UpdateUser) error
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
