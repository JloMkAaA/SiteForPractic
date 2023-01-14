package service

import "github.com/JloMkAaA/SiteForPractic/pkg/repository"

type ManipulateUser interface {
}

type Service struct {
	ManipulateUser
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
