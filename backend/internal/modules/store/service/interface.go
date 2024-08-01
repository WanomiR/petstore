package service

import (
	"backend/internal/modules/store/entities"
	"context"
)

type StoreServicer interface {
	CreateOrder(ctx context.Context, order entities.Order) (int, error)
	GetOrderById(ctx context.Context, orderId int) (entities.Order, error)
	DeleteOrder(ctx context.Context, orderId int) error
}
