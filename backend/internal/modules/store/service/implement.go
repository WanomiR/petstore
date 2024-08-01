package service

import (
	"backend/internal/modules/store/entities"
	"backend/internal/repository"
)

type StoreService struct {
	DB repository.Repository
}

func NewStoreService(db repository.Repository) *StoreService {
	return &StoreService{DB: db}
}

func (s *StoreService) CreateOrder(order entities.Order) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StoreService) GetOrderById(orderId int) (entities.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StoreService) DeleteOrder(orderId int) error {
	//TODO implement me
	panic("implement me")
}
