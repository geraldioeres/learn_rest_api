package users_test

import (
	"context"
	"learn_api/business/users"
	_mockUserRepository "learn_api/business/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockUserRepository.Repository

var userService users.Usecase
var userDomain users.Domain

func setup() {
	userService = users.NewUserUsecase(&userRepository, time.Hour*1)
	userDomain = users.Domain{
		Id:       1,
		Name:     "Alterra",
		Email:    "alterra@gmail.com",
		Password: "abc123",
		Address:  "Malang",
		Token:    "123",
	}
}

func TestLogin(t *testing.T) {
	setup()
	userRepository.On("Login",
		mock.Anything,
		mock.AnythingOfType("string"),
		mock.AnythingOfType("string")).Return(userDomain, nil).Once()

	t.Run("Test Case | Valid Login", func(t *testing.T) {
		user, err := userService.Login(context.Background(), users.Domain{
			Email:    "alterra@gmail.com",
			Password: "123",
		})

		assert.Nil(t, err)
		assert.Equal(t, "Alterra", user.Name)
	})
}
