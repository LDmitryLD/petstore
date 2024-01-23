package service

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/user/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	user     = models.User{ID: 1, Username: "username"}
	id       = 1
	username = "username"
)

func TestUserService_Create(t *testing.T) {
	storage := mocks.NewUserStorager(t)
	storage.On("Create", user).Return(id)

	userService := NewUserService(storage)

	got := userService.Create(user)

	assert.Equal(t, id, got)
}

func TestUserService_GetByUsername(t *testing.T) {
	storage := mocks.NewUserStorager(t)
	storage.On("GetByUsername", username).Return(user, nil)

	userService := NewUserService(storage)

	got, err := userService.GetByUsername(username)

	assert.Nil(t, err)
	assert.Equal(t, user, got)
}

func TestUserService_Update(t *testing.T) {
	storage := mocks.NewUserStorager(t)
	storage.On("Update", username, user).Return(nil)

	userService := NewUserService(storage)

	got := userService.Update(username, user)

	assert.Nil(t, got)
}

func TestUserService_Delete(t *testing.T) {
	storage := mocks.NewUserStorager(t)
	storage.On("Delete", username).Return(nil)

	userService := NewUserService(storage)

	got := userService.Delete(username)

	assert.Nil(t, got)
}
