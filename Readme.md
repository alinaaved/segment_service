# Segment Service

Простой REST API сервис на Go для управления пользовательскими сегментами (VK кейс).

## 📦 Возможности

- Создание и удаление сегментов
- Назначение сегментов пользователям (по ID или случайно по проценту)
- Получение списка сегментов пользователя по его ID
- HTML-интерфейс для ручной работы с API

## 🧱 Структура проекта

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
├── web/                   # HTML-интерфейс
│   └── index.html
├── docker-compose.yml     # контейнер для PostgreSQL
├── .env.example           # параметры подключения
├── go.mod
└── README.md
```

---

## 🚀 Быстрый старт

### 1. Запусти PostgreSQL через Docker

```bash
docker-compose up -d
```

### 2. Запусти backend-сервер

```bash
go run main.go
```

### 3. Открой HTML-интерфейс

Открой файл `frontend.html` (или `web/index.html`) в браузере. Он обращается к `localhost:8080`.

---

## 📌 API Endpoints

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

## 🧪 Тестирование

Юнит-тесты лежат в `handlers_test.go` и используют SQLite in-memory:
```bash
go test ./handlers
```

Покрытие:
- создание сегмента
- назначение пользователю
- проверка сегментов пользователя

---

## 💡 Полезные команды

```bash
# посмотреть таблицы внутри базы
\dt

# структура таблицы users
\d users

# данные
SELECT * FROM users;
SELECT * FROM segments;
SELECT * FROM user_segments;
```

Запускаются внутри контейнера:
```bash
docker exec -it segment_postgres psql -U postgres -d segment_service
```

---

## ✨ Пример curl-запросов

```bash
curl -X POST http://localhost:8080/segments \
     -H "Content-Type: application/json" \
     -d '{"name": "MAIL_GPT"}'

curl -X POST http://localhost:8080/segments/MAIL_GPT/assign \
     -H "Content-Type: application/json" \
     -d '{"user_ids": [15230, 19241]}'

curl http://localhost:8080/users/15230/segments
```

---

## 🛠 Зависимости

- Go 1.20+
- Gin
- GORM
- PostgreSQL (через Docker)
- SQLite (для тестов)

---

## 🧠 Автор/участник
Алина Ведерникова · 2025