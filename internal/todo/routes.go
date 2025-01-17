package todo

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, db *sql.DB) {
	handler := NewHandler(db)

	todoGroup := app.Group("/todos")
	todoGroup.Get("/", handler.GetAll)
	todoGroup.Post("/", handler.Create)
	todoGroup.Get("/:id", handler.GetByID)
	todoGroup.Put("/:id", handler.Update)
	todoGroup.Delete("/:id", handler.Delete)
}
