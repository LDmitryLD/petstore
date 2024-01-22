package service

import "projects/LDmitryLD/petstore/app/internal/models"

type Userer interface {
	Create(user models.User) int
	GetByUsername(username string) (models.User, error)
	Update(username string, user models.User) error
	Delete(username string) error
}
