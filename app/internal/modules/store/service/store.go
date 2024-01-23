package service

import (
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/store/storage"
)

type StoreService struct {
	storage storage.StoreStorager
}

func NewStoreService(storage storage.StoreStorager) *StoreService {
	return &StoreService{storage: storage}
}

func (s *StoreService) Create(order models.Order) models.Order {
	return s.storage.CreateOrder(order)
}

func (s *StoreService) Delete(id int) error {
	return s.storage.Delete(id)
}

func (s *StoreService) GetByID(id int) (models.Order, error) {
	return s.storage.GetByID(id)
}

func (s *StoreService) Inventory() map[string]interface{} {
	inventory := s.storage.Inventory()

	res := make(map[string]interface{})

	for status, val := range inventory {
		res[status] = val
	}

	return res
}
