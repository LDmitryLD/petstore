package service

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/user/storage"
)

type UserService struct {
	storage storage.UserStorager
}

func NewUserService(storage storage.UserStorager) *UserService {
	return &UserService{storage: storage}
}

func (u *UserService) Create(user models.User) int {
	return u.storage.Create(user)
}

func (u *UserService) GetByUsername(username string) (models.User, error) {
	return u.storage.GetByUsername(username)
}

func (u *UserService) Update(username string, user models.User) error {
	return u.storage.Update(username, user)
}

func (u *UserService) Delete(username string) error {
	return u.storage.Delete(username)
}
