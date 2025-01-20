CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,                    -- Уникальный идентификатор
    name VARCHAR(255) NOT NULL,               -- Название задачи
    description TEXT,                         -- Описание задачи (опционально)
    created_at TIMESTAMP DEFAULT NOW() NOT NULL, -- Дата создания
    updated_at TIMESTAMP DEFAULT NOW() NOT NULL, -- Дата обновления
    status VARCHAR(50) NOT NULL DEFAULT 'pending' -- Статус задачи
);
