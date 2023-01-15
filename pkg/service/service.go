package service

import (
	"github.com/JloMkAaA/SiteForPractic"
	"github.com/JloMkAaA/SiteForPractic/pkg/repository"
)

type ManipulateUser interface {
	CreateUser(user SiteForPractic.User) (int, error)
}

type Service struct {
	ManipulateUser
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
