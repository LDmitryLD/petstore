package storage

import "projects/LDmitryLD/petstore/app/internal/models"

type UserStorager interface {
	Create(user models.User) int
	GetByUsername(userName string) (models.User, error)
	Update(userName string, user models.User) error
	Delete(userName string) error
}
