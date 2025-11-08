# Candidate Management CRUD App

Простое CRUD приложение на Golang с ипользованием PostgreSQL и Docker Compose.
Позволяет создавать, просматривать, обновлять и удалять кандидатов в базе данных.

## Архитектура 
Приложение реализовано по принципам слоистой архитектуры (Layered Architecture) и чистого кода (Clean Code).
Каждый слой имеет своё назначение и не нарушает принцип разделения ответственности.

### Слои приложения
| Слой | Назначение |
|------|-------------|
| **Models** | Определяет структуры данных и их поля. Например, Candidate |
| **Repository** | Выполняет SQL-запросы и взаимодействует с базой данных. |
| **Services** | Содержит бизнес-логику и валидацию данных. |
| **Handlers** | Принимает HTTP-запросы и возвращает ответы клиенту. |

## Структура проекта:
```
final_project
├── cmd/app/main.go                          # Точка входа в программу
├── internal/                                # Внутренний код проекта
│   ├── config/config.go                     # Конфигурация приложения
│   ├── db/db.go                             # Работа с БД (connection pool)
│   ├── handlers/candidate_handlers.go       # HTTP-обработчики (слой Handler)
│   ├── services/candidate_sevice.go         # Бизнес-логика (слой Services)
│   ├── repository/candidate_repository.go   # Работа с данными (слой Repository)
│   └── models/candidate.go                  # Структуры данных
└── pkg/                                     # Переиспользуемые пакеты
│   └── http_utils/http.go                   # Утилиты для HTTP
├── Dockerfile
├── docker-compose.yml
└── README.md
``` 

## Эндпойнты:
### 1. Получить список всех кандидатов
**GET** `/candidates` — получить всех кандидатов
http://localhost:8080/candidates

### 2. Добавить нового кандидата 
**POST** `/candidates/create` — получить всех кандидатов
http://localhost:8080/candidates/create
Body → raw → JSON
```
{
  "name": "Aidana Yermukhambetova",
  "position": "Golang developer",
  "email": "aidana.yermukhambetova@gmail.com",
  "phone": "+12345678"
}

```

### 3. Обновить данные кандидата 
**PUT** `/candidates/update?id={id}` — получить всех кандидатов
http://localhost:8080/candidates/update?id={id}
Body → raw → JSON
```
{
  "name": "Aidana Yermukhambetova",
  "position": "Senior Golang developer",
  "email": "aidana.yermukhambetova@gmail.com",
  "phone": "+12345678"
}

```

### 4. Удалить кандидата 
**PUT** `/candidates/delete?id={id}` — получить всех кандидатов
http://localhost:8080/candidates/delete?id={id}

## Тесты 

