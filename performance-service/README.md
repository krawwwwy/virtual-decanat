# Performance Service

Микросервис для работы с успеваемостью студентов: расчёт рейтинга, посещаемость, оценки, долги.

## Архитектура
- Clean Architecture (cmd, internal, api, configs, test)
- REST API (chi)
- PostgreSQL
- OpenTelemetry

## Основные endpoints
- GET    /health
- GET    /performance/{student_id}
- GET    /performance/{student_id}/grades
- GET    /performance/{student_id}/attendance
- GET    /performance/{student_id}/debts
- GET    /performance/{student_id}/rating

## Фронтенд

- Тестовый SPA-фронт доступен по `/` и `/test.html`.
- Любой несуществующий путь (кроме API) отдаёт `frontend.html` (SPA fallback).
- Для теста API: открой http://localhost:8084/ в браузере.

## Docker

### Сборка и запуск

```sh
docker compose build performance-service
docker compose up -d performance-service
```

- `frontend.html` копируется в контейнер и доступен приложению.
- ENV переменные для подключения к БД задаются через docker-compose.

## Структура
- cmd/         — entrypoint
- internal/    — бизнес-логика, handler, service, repository, model
- configs/     — конфиги
- frontend.html — SPA-фронт для теста API

## ENV
- DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME

## Локальный запуск
```sh
go run ./cmd/main.go
``` 