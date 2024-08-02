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

	query := `SELECT id, username, first_name, last_name, email, password, phone, user_status 
			    FROM users 
				 WHERE username = $1 AND is_deleted = FALSE`

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

func (db *PostgresDBRepo) CreateUser(ctx context.Context, user entities.User) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `INSERT INTO users (username, first_name, last_name, email, password, phone, user_status, is_deleted) 
			  	 VALUES ($1, $2, $3, $4, $5, $6, $7, FALSE)
				 RETURNING id`

	var id int
	if err := db.conn.QueryRowContext(ctx, query, user.Username, user.FirstName, user.LastName,
		user.Email, user.Password, user.Phone, user.UserStatus).Scan(&id); err != nil {
		return 0, e.Wrap("failed to execute query", err)
	}

	return id, nil
}

func (db *PostgresDBRepo) DeleteUser(ctx context.Context, username string) error {
	ctx, cancel := context.WithTimeout(ctx, db.timeout)
	defer cancel()

	query := `UPDATE users SET is_deleted = TRUE WHERE username = $1`

	if _, err := db.conn.ExecContext(ctx, query, username); err != nil {
		return e.Wrap("failed to execute query", err)
	}

	return nil
}
