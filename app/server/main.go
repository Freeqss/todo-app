package main

import (
	"log"
	"todo-app/internal/database"
	"todo-app/internal/todo"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Инициализация базы данных
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Проверяем соединение с базой данных
	if err := db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}
	log.Println("Connected to the database successfully.")

	// Создаем приложение Fiber
	app := fiber.New()

	// Регистрация маршрутов
	todo.RegisterRoutes(app, db)

	// Запуск сервера
	log.Println("Starting server on :3000")
	log.Fatal(app.Listen(":3000"))
}
