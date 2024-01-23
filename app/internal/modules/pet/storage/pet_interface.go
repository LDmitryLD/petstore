package storage

import (
	"mime/multipart"
	"projects/LDmitryLD/petstore/app/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=PetStorager
type PetStorager interface {
	Create(pet models.Pet) models.Pet
	Update(pet models.Pet) error
	Delete(petID int) error
	GetByID(petID int) (models.Pet, error)
	GetList() []models.Pet
	UploadImage(file multipart.File, fileName string, id int) error
	FindByStatus(values []string) []models.Pet
	UpdateByID(petID int, name, status string) error
}
