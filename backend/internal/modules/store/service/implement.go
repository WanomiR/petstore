package service

import (
	"backend/internal/modules/store/entities"
	"backend/internal/repository"
	"context"
	"errors"
	"time"
)

type StoreService struct {
	DB repository.Repository
}

func NewStoreService(db repository.Repository) *StoreService {
	return &StoreService{DB: db}
}

func (s *StoreService) CreateOrder(ctx context.Context, order entities.Order) (int, error) {
	if order.PetId == 0 || order.Quantity == 0 {
		return 0, errors.New("order id and quantity must be greater than zero")
	}

	if order.ShipDate.IsZero() {
		order.ShipDate = time.Now()
	}

	if order.Status == "" {
		order.Status = "placed"
	}

	if !s.statusIsValid(order.Status) {
		return 0, errors.New("invalid order status")
	}

	orderId, err := s.DB.CreateOrder(ctx, order)
	if err != nil {
		return 0, err
	}

	return orderId, nil
}

func (s *StoreService) GetOrderById(ctx context.Context, orderId int) (entities.Order, error) {
	return s.DB.GetOrderById(ctx, orderId)
}

func (s *StoreService) DeleteOrder(ctx context.Context, orderId int) error {
	return s.DB.DeleteOrder(ctx, orderId)
}

func (s *StoreService) statusIsValid(status string) bool {
	return status == "placed" || status == "approved" || status == "delivered"
}
