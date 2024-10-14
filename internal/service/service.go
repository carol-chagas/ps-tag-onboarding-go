package service

import (
	"context"
	"errors"
	"regexp"

	"ps-tag-onboarding-go/internal/domain"
	"ps-tag-onboarding-go/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) validateUser(user domain.User) error {
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(user.Email) {
		return errors.New("invalid email")
	}

	if user.Age < 18 {
		return errors.New("invalid age")
	}

	return nil
}

func (s *UserService) SaveUser(ctx context.Context, user domain.User) error {
	if err := s.validateUser(user); err != nil {
		return err
	}

	existingUser, _ := s.repo.FindByEmail(ctx, user.Email)
	if existingUser.ID != "" && existingUser.ID != user.ID {
		return errors.New("email already exists")
	}

	return s.repo.Save(ctx, user)
}

func (s *UserService) UpdateUser(ctx context.Context, user domain.User) error {
	if err := s.validateUser(user); err != nil {
		return err
	}
	return s.repo.Update(ctx, user)
}

func (s *UserService) FindUserByID(ctx context.Context, id string) (domain.User, bool) {
	return s.repo.FindByID(ctx, id)
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	user, found := s.repo.FindByID(ctx, id)
	if !found {
		return errors.New("user not found")
	}
	return s.repo.Delete(ctx, user.ID)
}
