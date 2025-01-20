package todo

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	db *sql.DB
}

// Конструктор хендлера
func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	// Ответ, подтверждающий, что сервер работает
	return c.JSON(fiber.Map{
		"message": "Server is running",
	})
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	// Запрос к базе данных для получения всех задач
	rows, err := h.db.Query("SELECT id, name, description, created_at, updated_at, status FROM tasks")
	if err != nil {
		log.Printf("Error querying database: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve tasks",
		})
	}
	defer rows.Close()

	var tasks []Task
	// Проходим по строкам результата
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.CreatedAt, &task.UpdatedAt, &task.Status); err != nil {
			log.Printf("Error scanning row: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to parse task data",
			})
		}
		tasks = append(tasks, task)
	}

	// Проверяем, есть ли ошибки при обработке строк
	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve tasks",
		})
	}

	// Возвращаем все задачи в формате JSON
	return c.JSON(tasks)
}

// Хендлер для получения всех задач
func (h *Handler) Create(c *fiber.Ctx) error {
	// Структура для данных задачи
	type CreateTaskRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	// Парсим тело запроса
	var request CreateTaskRequest
	if err := c.BodyParser(&request); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request data",
		})
	}

	// Вставляем данные в базу
	_, err := h.db.Exec(
		"INSERT INTO tasks (name, description, status) VALUES ($1, $2, $3)",
		request.Name, request.Description, request.Status,
	)
	if err != nil {
		log.Printf("Error inserting task: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create task",
		})
	}

	// Возвращаем сообщение об успешном создании
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Task created successfully",
	})
}
