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

## Структура
- cmd/         — entrypoint
- internal/    — бизнес-логика, usecase, repo
- api/         — хендлеры, контракты
- configs/     — конфиги
- test/        — тесты

## ENV
- DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME

## Запуск
```
go run ./cmd/main.go
``` 