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
RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway ./cmd/main.go

# Используем минимальный образ для запуска приложения
FROM alpine:latest

# Устанавливаем timezone
RUN apk --no-cache add tzdata
ENV TZ=Europe/Moscow

# Добавляем сертификаты для HTTPS запросов к другим сервисам
RUN apk --no-cache add ca-certificates

# Копируем исполняемый файл из предыдущего этапа
COPY --from=builder /app/api-gateway /api-gateway

# Создаем непривилегированного пользователя
RUN adduser -D -g '' appuser
USER appuser

# Запускаем приложение
ENTRYPOINT ["/api-gateway"] 