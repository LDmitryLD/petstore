package service

import (
	"mime/multipart"
	"os"
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/pet/storage/mocks"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPetService_Create(t *testing.T) {
	pet := models.Pet{Name: "Alma"}
	expect := CreateOut{Pet: pet}

	storageMock := mocks.NewPetStorager(t)
	storageMock.On("Create", pet).Return(pet)

	petService := NewPetService(storageMock)

	got := petService.Create(CreateIn{Pet: pet})

	assert.Equal(t, expect, got)
}

func TestPetService_Update(t *testing.T) {
	pet := models.Pet{Name: "Alma"}
	expect := UpdateOut{Error: nil}

	storageMock := mocks.NewPetStorager(t)
	storageMock.On("Update", pet).Return(nil)

	petService := NewPetService(storageMock)

	got := petService.Update(UpdateIn{Pet: pet})

	assert.Equal(t, expect, got)
}

func TestPetService_Delete(t *testing.T) {
	petID := 1

	storageMock := mocks.NewPetStorager(t)
	storageMock.On("Delete", petID).Return(nil)

	petService := NewPetService(storageMock)

	got := petService.Delete(petID)

	assert.Nil(t, got)
}

func TestPetService_GetByID(t *testing.T) {
	pet := models.Pet{ID: 1, Name: "Alma"}
	petID := 1
	expect := GetByIDOut{Pet: pet}

	storageMock := mocks.NewPetStorager(t)
	storageMock.On("GetByID", petID).Return(pet, nil)

	petService := NewPetService(storageMock)

	got := petService.GetByID(GetByIDIn{PetID: petID})

	assert.Equal(t, expect, got)
}

func TestPetService_GetList(t *testing.T) {
	pets := []models.Pet{{ID: 3}, {ID: 2}, {ID: 1}}

	storageMock := mocks.NewPetStorager(t)
	storageMock.On("GetList").Return(pets)

	petService := NewPetService(storageMock)

	got := petService.GetList()

	sort.Slice(pets, func(i, j int) bool {
		return i < j
	})
	expect := GetListOut{List: pets}

	assert.Equal(t, expect, got)
}

func TestPetService_UploadImage(t *testing.T) {
	fileName := "filename"
	petID := 1
	var ptr uintptr
	file := os.NewFile(ptr, "filename")

	serviceMock := mocks.NewPetStorager(t)
	serviceMock.On("UploadImage", file, fileName, petID).Return(nil)

	petService := NewPetService(serviceMock)

	got := petService.UploadImage(UploadIn{PetID: petID, File: file, Handler: &multipart.FileHeader{Filename: "filename"}})

	assert.Nil(t, got.Error)
	assert.Equal(t, fileName, got.FileName)
}

func TestPetService_FindByStatus(t *testing.T) {
	pets := []models.Pet{{ID: 3, Status: "status1"}, {ID: 2, Status: "status1"}, {ID: 1, Status: "status3"}}
	values := []string{"status1"}

	storageMock := mocks.NewPetStorager(t)
	storageMock.On("FindByStatus", values).Return(pets[:2])

	petService := NewPetService(storageMock)

	got := petService.FindByStatus(values)

	assert.Equal(t, pets[:2], got)
}

func TestPetService_UpdateByID(t *testing.T) {
	in := UpdateByIDIn{
		PetID:  1,
		Name:   "NewName",
		Status: "NewStatus",
	}

	storageMock := mocks.NewPetStorager(t)
	storageMock.On("UpdateByID", in.PetID, in.Name, in.Status).Return(nil)

	petService := NewPetService(storageMock)

	got := petService.UpdateByID(in)

	assert.Nil(t, got)
}
