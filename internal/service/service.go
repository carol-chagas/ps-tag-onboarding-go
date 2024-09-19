package service

import (
	"ps-tag-onboarding-go/internal/domain"
	"ps-tag-onboarding-go/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) SaveUser(user domain.User) error {
	return s.repo.Save(user)
}

func (s *UserService) UpdateUser(user domain.User) error {
	return s.repo.Update(user)
}

func (s *UserService) FindUserByID(id string) (domain.User, bool) {
	return s.repo.FindByID(id)
}
