package service

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/store/storage/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreService_Create(t *testing.T) {
	order := models.Order{ID: 1}

	storageMock := mocks.NewStoreStorager(t)
	storageMock.On("CreateOrder", order).Return(order)

	service := NewStoreService(storageMock)

	got := service.Create(order)

	assert.Equal(t, order, got)
}

func TestStoreService_Delete(t *testing.T) {
	id := 1

	storageMock := mocks.NewStoreStorager(t)
	storageMock.On("Delete", id).Return(nil)

	service := NewStoreService(storageMock)

	got := service.Delete(id)

	assert.Nil(t, got)
}

func TestStorageService_GetByID(t *testing.T) {
	id := 1
	order := models.Order{ID: 1}

	storageMock := mocks.NewStoreStorager(t)
	storageMock.On("GetByID", id).Return(order, nil)

	service := NewStoreService(storageMock)

	got, err := service.GetByID(id)

	assert.Nil(t, err)
	assert.Equal(t, order, got)
}

func TestStorageService_Inventory(t *testing.T) {
	invent := map[string]int{"status": 10}
	expect := map[string]interface{}{"status": 10}

	storageMock := mocks.NewStoreStorager(t)
	storageMock.On("Inventory").Return(invent)

	service := NewStoreService(storageMock)

	got := service.Inventory()

	assert.Equal(t, expect, got)
}
