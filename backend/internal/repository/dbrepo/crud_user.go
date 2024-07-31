package dbrepo

import (
	"backend/internal/lib/e"
	"backend/internal/modules/user/entities"
	"context"
	"database/sql"
	"errors"
)

func (db *PostgresDBRepo) GetUserByUsername(ctx context.Context, username string) (entities.User, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `SELECT id, username, first_name, last_name, email, password, phone, user_status FROM users WHERE username = $1`

	var user entities.User
	err := db.conn.QueryRowContext(ctx, query, username).Scan(
		&user.Id,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.UserStatus,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return entities.User{}, errors.New("user not found")
	} else if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

func (db *PostgresDBRepo) UpdateUser(ctx context.Context, user entities.User) error {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4, phone = $5, user_status = $6 WHERE username = $7`

	if _, err := db.conn.ExecContext(ctx, query, user.FirstName, user.LastName,
		user.Email, user.Password, user.Phone, user.UserStatus, user.Username); err != nil {
		return e.Wrap("failed to execute query", err)
	}

	return nil
}
