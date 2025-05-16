# Используем образ golang для сборки приложения
FROM golang:1.21-alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы модулей и скачиваем зависимости
COPY go.mod go.sum* ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/main.go

# Используем минимальный образ для запуска приложения
FROM alpine:latest

# Устанавливаем timezone
RUN apk --no-cache add tzdata
ENV TZ=Europe/Moscow

# Добавляем сертификаты для HTTPS запросов
RUN apk --no-cache add ca-certificates

# Копируем исполняемый файл из предыдущего этапа
COPY --from=builder /app/auth-service /auth-service

# Создаем непривилегированного пользователя
RUN adduser -D -g '' appuser
USER appuser

# Определяем переменные окружения по умолчанию
ENV DB_HOST=auth-db
ENV DB_PORT=5432
ENV DB_NAME=auth_service

# Запускаем приложение
ENTRYPOINT ["/auth-service"]
