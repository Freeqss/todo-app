package main

import (
	"database/sql"
	"log"
	"todo-app/internal/todo"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	// Подключение к базе данных
	db, err := sql.Open("postgres", "postgres://postgres:12345@localhost:1234/tododb?sslmode=disable")
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Создание приложения Fiber
	app := fiber.New()

	// Регистрация маршрутов
	todo.RegisterRoutes(app, db)

	// Запуск приложения на порту 3000
	log.Fatal(app.Listen(":3000"))
}
