package storage

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	id     = 1
	name   = "Valan"
	status = "status"
	pet    = models.Pet{
		ID:     1,
		Name:   "Valan",
		Status: "status",
	}
)

func TestPetStorage_Create(t *testing.T) {

	storage := NewPetStorage()

	res := storage.Create(pet)

	assert.Equal(t, pet, res)
	assert.Equal(t, &pet, storage.primaryKeyIDx[id])
}

func TestPetStorage_Update_NotFound(t *testing.T) {
	storage := NewPetStorage()

	res := storage.Update(pet)

	assert.NotNil(t, res)
}

func TestPetStorage_Update(t *testing.T) {
	storage := NewPetStorage()
	storage.primaryKeyIDx[id] = &pet

	res := storage.Update(pet)

	assert.Nil(t, res)
}

func TestPetStorage_Delete_NotFound(t *testing.T) {
	storage := NewPetStorage()

	res := storage.Delete(id)

	assert.NotNil(t, res)
}

func TestPetStorage_Delete(t *testing.T) {
	storage := NewPetStorage()
	storage.primaryKeyIDx[id] = &pet

	res := storage.Delete(id)

	assert.Nil(t, res)
}

func TestPetStorage_GetByID_NotFound(t *testing.T) {
	storage := NewPetStorage()

	res, err := storage.GetByID(id)

	assert.Equal(t, models.Pet{}, res)
	assert.NotNil(t, err)
}

func TestPetStorage_GetByID(t *testing.T) {
	storage := NewPetStorage()
	storage.primaryKeyIDx[id] = &pet

	res, err := storage.GetByID(id)

	assert.Equal(t, pet, res)
	assert.Nil(t, err)
}

func TestPetStorage_GetList_Empty(t *testing.T) {
	storage := NewPetStorage()

	res := storage.GetList()

	assert.Equal(t, []models.Pet{}, res)
}

func TestPetStorage_GetList(t *testing.T) {
	expect := []models.Pet{pet}

	storage := NewPetStorage()
	storage.data = append(storage.data, &pet)
	res := storage.GetList()

	assert.Equal(t, expect, res)
}

func TestPetStorage_FindByStatus(t *testing.T) {
	storage := NewPetStorage()
	storage.data = append(storage.data, &pet)
	res := storage.FindByStatus([]string{"status"})

	assert.Equal(t, []models.Pet{pet}, res)
}

func TestPetStorage_UpdateByID_NotFound(t *testing.T) {
	storage := NewPetStorage()

	err := storage.UpdateByID(id, name, status)

	assert.NotNil(t, err)
}

func TestPetStorage_UpdateByID(t *testing.T) {
	storage := NewPetStorage()
	storage.data = append(storage.data, &pet)
	storage.primaryKeyIDx[id] = &pet

	err := storage.UpdateByID(id, name, status)

	assert.Nil(t, err)
}
