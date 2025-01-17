package database

import (
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

func Connect() (*sql.DB, error) {
	// Настройте DSN в соответствии с вашим окружением
	dsn := "user=postgres password=yourpassword dbname=tododb sslmode=disable"
	return sql.Open("postgres", dsn)
}
