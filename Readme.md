# Segment Service

Простой REST API сервис на Go для управления пользовательскими сегментами (VK кейс).

## Возможности

- Создание и удаление сегментов
- Назначение сегментов пользователям (по ID или случайно по проценту)
- Получение списка сегментов пользователя по его ID

## Структура проекта

```
segment_service/
├── cmd/                    # точка входа (main.go)
│   └── server/
│       └── main.go
├── internal/
│   ├── db/                # инициализация подключения к БД
│   │   └── db.go
│   ├── models/            # модели User и Segment
│   │   └── models.go
│   ├── handlers/          # обработчики HTTP запросов
│   │   ├── handlers.go
│   │   └── handlers_test.go
│   └── routes/            # маршруты для API
│       └── routes.go
├── docker-compose.yml     # контейнер для PostgreSQL
├── go.mod
└── README.md
```

---

## Инструкция

### 1. Запусти PostgreSQL через Docker

```bash
docker-compose up -d
```

### 2. Запусти backend-сервер

`go run main.go`

---

## API Endpoints

### POST /segments
Создание нового сегмента:
```json
{
  "name": "MAIL_GPT"
}
```

### DELETE /segments/:name
Удаление сегмента

### POST /segments/:name/assign
Назначение сегмента:
- по ID:
```json
{
  "user_ids": [1, 2, 3]
}
```
- по проценту:
```json
{
  "percent": 30
}
```

### GET /users/:id/segments
Получение списка сегментов пользователя

---

## Тестирование

Юнит-тесты лежат в `handlers_test.go` и используют SQLite in-memory:
```bash
go test ./handlers
```

Покрытие:
- создание сегмента
- назначение пользователю
- проверка сегментов пользователя

---

## Пример curl-запросов

```
curl -X POST http://localhost:8080/segments \
     -H "Content-Type: application/json" \
     -d '{"name": "MAIL_GPT"}'

curl -X POST http://localhost:8080/segments/MAIL_GPT/assign \
     -H "Content-Type: application/json" \
     -d '{"user_ids": [15230, 19241]}'

curl http://localhost:8080/users/15230/segments
```

---

## Зависимости

- Go 1.20+
- Gin
- GORM
- PostgreSQL (через Docker)
- SQLite (для тестов)

---

## Автор
Алина Ведерникова · 2025
