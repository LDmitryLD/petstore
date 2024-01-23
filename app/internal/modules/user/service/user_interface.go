package service

import "projects/LDmitryLD/petstore/app/internal/models"

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=Userer
type Userer interface {
	Create(user models.User) int
	GetByUsername(username string) (models.User, error)
	Update(username string, user models.User) error
	Delete(username string) error
}
