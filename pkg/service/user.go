package service

import (
	"github.com/JloMkAaA/SiteForPractic"
	"github.com/JloMkAaA/SiteForPractic/pkg/repository"
)

type UserService struct {
	repo repository.UserManipulated
}

func NewUserService(repo repository.UserManipulated) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserById(userId int) (SiteForPractic.User, error) {
	return s.repo.GetUserById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *UserService) Update(userId int, input SiteForPractic.UpdateUser) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, input)
}
