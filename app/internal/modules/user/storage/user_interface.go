package storage

import "projects/LDmitryLD/petstore/app/internal/models"

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=UserStorager
type UserStorager interface {
	Create(user models.User) int
	GetByUsername(userName string) (models.User, error)
	Update(userName string, user models.User) error
	Delete(userName string) error
}
