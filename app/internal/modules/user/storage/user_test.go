package storage

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	id       = 1
	username = "username"
	user     = models.User{
		ID:       1,
		Username: "username",
	}
)

func TestUserStorage_Create(t *testing.T) {
	storage := NewUserStorage()

	got := storage.Create(user)

	assert.Equal(t, id, got)
}

func TestUserStorage_GetByUsername_NotFound(t *testing.T) {
	storage := NewUserStorage()

	got, err := storage.GetByUsername(username)

	assert.Equal(t, models.User{}, got)
	assert.NotNil(t, err)
}

func TestUserStorage_GetByUsername(t *testing.T) {
	storage := NewUserStorage()
	storage.userNameKeys[username] = &user

	got, err := storage.GetByUsername(username)

	assert.Equal(t, user, got)
	assert.Nil(t, err)
}

func TestUserStorage_Update_NotFound(t *testing.T) {
	storage := NewUserStorage()

	err := storage.Update(username, user)

	assert.NotNil(t, err)
}

func TestUserStorage_Update(t *testing.T) {
	storage := NewUserStorage()
	storage.userNameKeys[username] = &user

	err := storage.Update(username, user)

	assert.Nil(t, err)
}

func TestUserStorage_Delete_NotFound(t *testing.T) {
	storage := NewUserStorage()

	err := storage.Delete(username)

	assert.NotNil(t, err)
}

func TestUserStorage_Delete(t *testing.T) {
	storage := NewUserStorage()
	storage.userNameKeys[username] = &user
	storage.users = append(storage.users, &user)

	err := storage.Delete(username)

	assert.Nil(t, err)
}
