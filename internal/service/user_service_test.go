package service

import (
	"testing"

	"ps-tag-onboarding-go/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

// Delete implements repository.UserRepository.
func (m *MockUserRepository) Delete(id string) error {
	panic("unimplemented")
}

// FindAll implements repository.UserRepository.
func (m *MockUserRepository) FindAll() ([]domain.User, error) {
	panic("unimplemented")
}

func (m *MockUserRepository) Save(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByID(id string) (domain.User, bool) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Bool(1)
}

func (m *MockUserRepository) FindByName(firstName, lastName string) (domain.User, bool) {
	args := m.Called(firstName, lastName)
	return args.Get(0).(domain.User), args.Bool(1)
}

func Test_SaveAndValidateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	mockRepo.On("Save", mock.Anything).Return(nil)

	userService := NewUserService(mockRepo)

	user := domain.User{
		FirstName: "Ana",
		LastName:  "Chagas",
		Email:     "ana.chagas@example.com",
		Age:       23,
	}

	err := userService.SaveUser(user)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func Test_SaveUserInvalidEmail(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	user := domain.User{
		FirstName: "Ana",
		LastName:  "Chagas",
		Email:     "ana.chagas",
		Age:       23,
	}

	err := userService.SaveUser(user)

	assert.Error(t, err)
	mockRepo.AssertNotCalled(t, "Save")
}

func Test_SaveUserInvalidAge(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userService := NewUserService(mockRepo)

	user := domain.User{
		FirstName: "Ana",
		LastName:  "Chagas",
		Email:     "ana.chagas@example.com",
		Age:       17,
	}

	err := userService.SaveUser(user)
	assert.Error(t, err)
}
