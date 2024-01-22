package service

import (
	"mime/multipart"
	"projects/LDmitryLD/petstore/app/internal/models"
)

type Peterer interface {
	Create(in CreateIn) CreateOut
	Update(in UpdateIn) UpdateOut
	Delete(petID int) error
	GetByID(in GetByIDIn) GetByIDOut
	GetList() GetListOut
	UploadImage(in UploadIn) UploadOut
	FindByStatus(values []string) []models.Pet
	UpdateByID(in UpdateByIDIn) error
}

type CreateIn struct {
	Pet models.Pet
}

type CreateOut struct {
	Pet models.Pet
}

type UpdateIn struct {
	Pet models.Pet
}

type UpdateOut struct {
	Error error
}

type GetByIDIn struct {
	PetID int
}

type GetByIDOut struct {
	Error error
	Pet   models.Pet
}

type GetListOut struct {
	List []models.Pet
}

type UploadIn struct {
	PetID   int
	File    multipart.File
	Handler *multipart.FileHeader
}

type UploadOut struct {
	FileName string
	FileType string
	Error    error
}

type UpdateByIDIn struct {
	PetID  int
	Name   string
	Status string
}
