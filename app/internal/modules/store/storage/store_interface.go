package storage

import (
	"projects/LDmitryLD/petstore/app/internal/models"
)

//go:generate go run github.com/vektra/mockery/v2@v2.35.4 --name=StoreStorager
type StoreStorager interface {
	CreateOrder(order models.Order) models.Order
	Delete(id int) error
	GetByID(id int) (models.Order, error)
	Inventory() map[string]int
}
