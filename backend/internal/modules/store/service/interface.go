package service

import "backend/internal/modules/store/entities"

type StoreServicer interface {
	CreateOrder(order entities.Order) (int, error)
	GetOrderById(orderId int) (entities.Order, error)
	DeleteOrder(orderId int) error
}
