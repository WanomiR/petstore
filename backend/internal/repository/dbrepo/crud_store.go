package dbrepo

import (
	"backend/internal/modules/store/entities"
	"context"
)

func (db *PostgresDBRepo) GetOrderById(ctx context.Context, orderId int) (entities.Order, error) {
	panic("implement me")
}

func (db *PostgresDBRepo) CreateOrder(ctx context.Context, order entities.Order) (int, error) {
	panic("implement me")
}

func (db *PostgresDBRepo) DeleteOrder(ctx context.Context, orderId int) error {
	panic("implement me")
}
