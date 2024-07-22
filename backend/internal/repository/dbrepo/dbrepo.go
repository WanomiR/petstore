package dbrepo

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	conn    *sql.DB
	timeout time.Duration
}

func WithDBTimeout(timeout time.Duration) func(*PostgresDBRepo) {
	return func(db *PostgresDBRepo) {
		db.timeout = timeout
	}

}

func NewPostgresDBRepo(conn *sql.DB, options ...func(repo *PostgresDBRepo)) *PostgresDBRepo {
	db := &PostgresDBRepo{
		conn:    conn,
		timeout: time.Second * 3, // default timout
	}

	for _, option := range options {
		option(db)
	}

	return db
}

func (db *PostgresDBRepo) Connection() *sql.DB {
	return db.conn
}
