package service

import (
	"context"
	"testing"

	"ps-tag-onboarding-go/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (domain.User, bool) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.User), args.Bool(1)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (domain.User, bool) {
	args := m.Called(ctx, email)
	return args.Get(0).(domain.User), args.Bool(1)
}

func (m *MockUserRepository) FindByName(ctx context.Context, firstName, lastName string) (domain.User, bool) {
	args := m.Called(ctx, firstName, lastName)
	return args.Get(0).(domain.User), args.Bool(1)
}

func (m *MockUserRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
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

	ctx := context.Background()
	err := userService.SaveUser(ctx, user)

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

	ctx := context.Background()
	err := userService.SaveUser(ctx, user)

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

	ctx := context.Background()
	err := userService.SaveUser(ctx, user)
	assert.Error(t, err)
}
