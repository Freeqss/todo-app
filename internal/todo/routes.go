package todo

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, db *sql.DB) {
	handler := NewHandler(db)

	todoGroup := app.Group("/tasks")
	todoGroup.Get("/", handler.GetAll)
	todoGroup.Post("/", handler.Create)

	// Новый маршрут для проверки состояния сервера
	app.Get("/", handler.HealthCheck)
}
