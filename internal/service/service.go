package service

import (
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

func (s *UserService) SaveUser(user domain.User) error {
	if err := s.validateUser(user); err != nil {
		return err
	}

	existingUser, _ := s.repo.FindByName(user.FirstName, user.LastName)
	if existingUser.ID != "" && existingUser.ID != user.ID {
		return errors.New("user already exists")
	}

	return s.repo.Save(user)
}

func (s *UserService) FindUserByID(id string) (domain.User, bool) {
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) DeleteUser(id string) error {
	user, found := s.repo.FindByID(id)
	if !found {
		return errors.New("user not found")
	}
	return s.repo.Delete(user.ID)
}

func (s *UserService) UpdateUser(user domain.User) error {
	if err := s.repo.Update(user); err != nil {
		return err
	}
	return nil
}
