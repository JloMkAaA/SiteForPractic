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

func (s *UserService) GetUserById(userIdAuth, userId int) (SiteForPractic.User, error) {
	return s.repo.GetUserById(userIdAuth, userId)
}

func (s *UserService) Delete(userIdAuth, userId int) error {
	return s.repo.Delete(userIdAuth, userId)
}

func (s *UserService) Update(userIdAuth, userId int, input SiteForPractic.UpdateUser) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userIdAuth, userId, input)
}
