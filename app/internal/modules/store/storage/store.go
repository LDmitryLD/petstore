package storage

import (
	"fmt"
	"projects/LDmitryLD/petstore/app/internal/models"
	"sync"
)

var (
	ErrorNotFound = "not found"
)

type StoreStorage struct {
	orders             map[int]*models.Order
	inventoty          map[string]int
	autoIncrementCount int
	sync.Mutex
}

func NewStorage() *StoreStorage {
	return &StoreStorage{
		orders:             make(map[int]*models.Order, 13),
		inventoty:          make(map[string]int, 13),
		autoIncrementCount: 1,
	}
}

func (s *StoreStorage) CreateOrder(order models.Order) models.Order {
	s.Lock()
	defer s.Unlock()

	order.ID = s.autoIncrementCount
	s.autoIncrementCount++
	s.orders[order.ID] = &order
	s.inventoty[order.Status]++

	return order
}

func (s *StoreStorage) Delete(id int) error {
	s.Lock()
	defer s.Unlock()

	order, ok := s.orders[id]
	if !ok {
		return fmt.Errorf(ErrorNotFound)
	}

	if _, ok := s.inventoty[order.Status]; ok {
		s.inventoty[order.Status]--
	}

	delete(s.orders, id)

	return nil
}

func (s *StoreStorage) GetByID(id int) (models.Order, error) {
	s.Lock()
	defer s.Unlock()

	if order, ok := s.orders[id]; ok {
		return *order, nil
	}

	return models.Order{}, fmt.Errorf(ErrorNotFound)
}

func (s *StoreStorage) Inventory() map[string]int {
	s.Lock()
	defer s.Unlock()

	return s.inventoty
}
