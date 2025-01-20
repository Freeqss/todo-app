package database

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

func Connect() (*sql.DB, error) {
	dsn := "postgres://postgres:12345@localhost:1234/tododb?sslmode=disable"

	return sql.Open("postgres", dsn)
}
