package storage

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	id    = 1
	order = models.Order{
		ID:     1,
		Status: "status",
	}
)

func TestStoreStorage_Create(t *testing.T) {
	storage := NewStorage()

	res := storage.CreateOrder(order)

	assert.Equal(t, order, res)
}

func TestStoreStorage_Delete_NotFound(t *testing.T) {
	storage := NewStorage()

	err := storage.Delete(id)

	assert.NotNil(t, err)
}

func TestStoreStorage_Delete(t *testing.T) {
	storage := NewStorage()
	storage.orders[id] = &order
	storage.inventoty["status"]++

	err := storage.Delete(id)

	assert.Nil(t, err)
}

func TestStoreStorage_GetByID_NotFound(t *testing.T) {
	storage := NewStorage()

	res, err := storage.GetByID(id)

	assert.Equal(t, models.Order{}, res)
	assert.NotNil(t, err)
}

func TestStoreStorage_GetByID(t *testing.T) {
	storage := NewStorage()
	storage.orders[id] = &order

	res, err := storage.GetByID(id)

	assert.Equal(t, order, res)
	assert.Nil(t, err)
}

func TestStoreStorage_Inventory(t *testing.T) {
	expect := map[string]int{"status": 1}

	storage := NewStorage()
	storage.inventoty["status"]++

	res := storage.Inventory()

	assert.Equal(t, expect, res)
}
