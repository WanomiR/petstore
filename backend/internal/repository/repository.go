package repository

import "database/sql"

type Repository interface {
	Connection() *sql.DB
}
