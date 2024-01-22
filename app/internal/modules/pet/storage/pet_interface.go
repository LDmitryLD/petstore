package storage

import (
	"projects/LDmitryLD/petstore/app/internal/models"
)

type PetStorager interface {
	Create(pet models.Pet) models.Pet
	Update(pet models.Pet) error
	Delete(petID int) error
	GetByID(petID int) (models.Pet, error)
	GetList() []models.Pet
	UploadImage(fileName string, id int) error
	FindByStatus(values []string) []models.Pet
	UpdateByID(petID int, name, status string) error
}
