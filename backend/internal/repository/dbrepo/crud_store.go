package dbrepo

import (
	"backend/internal/lib/e"
	"backend/internal/modules/store/entities"
	"context"
	"database/sql"
	"errors"
)

func (db *PostgresDBRepo) GetOrderById(ctx context.Context, orderId int) (entities.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `SELECT id, pet_id, quantity, ship_date, status, is_complete 
				 FROM store 
				 WHERE id = $1 AND is_deleted = FALSE`

	var order entities.Order
	err := db.conn.QueryRowContext(ctx, query, orderId).Scan(&order.Id, &order.PetId, &order.Quantity,
		&order.ShipDate, &order.Status, &order.IsComplete)
	if errors.Is(err, sql.ErrNoRows) {
		return entities.Order{}, errors.New("order not found")
	} else if err != nil {
		return entities.Order{}, err
	}

	return order, nil
}

func (db *PostgresDBRepo) CreateOrder(ctx context.Context, order entities.Order) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `INSERT INTO store (pet_id, quantity, ship_date, status, is_complete, is_deleted)
				 VALUES ($1, $2, $3, $4, $5, FALSE) RETURNING id;`

	var orderId int
	if err := db.conn.QueryRowContext(ctx, query, order.PetId, order.Quantity,
		order.ShipDate, order.Status, order.IsComplete).Scan(&orderId); err != nil {
		return 0, e.Wrap("failed to execute query", err)
	}

	return orderId, nil
}

func (db *PostgresDBRepo) DeleteOrder(ctx context.Context, orderId int) error {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `UPDATE store SET is_deleted = TRUE WHERE id = $1`

	if _, err := db.conn.ExecContext(ctx, query, orderId); err != nil {
		return e.Wrap("failed to execute query", err)
	}

	return nil
}
