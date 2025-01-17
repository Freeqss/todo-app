package todo

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	// Логика для получения всех задач
	return c.JSON([]string{"TODO"})
}

func (h *Handler) Create(c *fiber.Ctx) error {
	// Логика для создания задачи
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Created"})
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	// Логика для получения задачи по ID
	return c.JSON(fiber.Map{"id": c.Params("id")})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	// Логика для обновления задачи
	return c.JSON(fiber.Map{"message": "Updated"})
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	// Логика для удаления задачи
	return c.SendStatus(fiber.StatusNoContent)
}
