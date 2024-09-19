package repository

import (
	"errors"
	"fmt"

	"ps-tag-onboarding-go/internal/domain"
)

type UserRepository struct {
	users map[string]domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[string]domain.User),
	}
}

func (r *UserRepository) Save(user domain.User) error {
	if user.FirstName == "" || user.LastName == "" || user.Email == "" {
		return errors.New("missing required fields")
	}
	if _, exists := r.FindByName(user.FirstName, user.LastName); exists {
		return fmt.Errorf("user %s %s already exists", user.FirstName, user.LastName)
	}
	if user.ID == "" {
		user.ID = fmt.Sprintf("%d", len(r.users)+1)
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) Update(user domain.User) error {
	if _, exists := r.users[user.ID]; !exists {
		return fmt.Errorf("user %s not found", user.ID)
	}

	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByID(id string) (domain.User, bool) {
	user, exists := r.users[id]
	return user, exists
}

func (r *UserRepository) FindByName(firstName, lastName string) (domain.User, bool) {
	for _, user := range r.users {
		if user.FirstName == firstName && user.LastName == lastName {
			return user, true
		}
	}
	return domain.User{}, false
}
