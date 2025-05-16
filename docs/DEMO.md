# Демонстрация работы сервиса аутентификации

В этом документе описаны способы демонстрации работы сервиса аутентификации в проекте "Виртуальный деканат".

## Предварительные требования

1. Установленный Go 1.21 или новее
2. Docker (для запуска базы данных PostgreSQL)
3. Доступ к интернету для скачивания зависимостей

## Способ 1: Запуск через Docker Compose

Самый простой способ запустить все сервисы через Docker Compose:

```bash
# Создать файл .env с необходимыми переменными окружения
cp .env.example .env

# Собрать и запустить все контейнеры
docker-compose -f decanat-dev-environment/docker-compose.yml up -d
```

## Способ 2: Запуск сервиса аутентификации локально

Для запуска только сервиса аутентификации:

### Linux/macOS

```bash
# Запустить PostgreSQL в Docker
docker run --name auth-db -e POSTGRES_USER=decanat_user -e POSTGRES_PASSWORD=decanat_password -e POSTGRES_DB=auth_service -p 5432:5432 -d postgres:latest

# Инициализировать БД
cat auth-service/scripts/init_db.sql | docker exec -i auth-db psql -U decanat_user -d auth_service

# Задать необходимые переменные окружения
export JWT_SECRET="virtual_decanat_secret_key_for_development"
export JWT_EXPIRATION="24h"
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_USER="decanat_user"
export DB_PASSWORD="decanat_password"
export DB_NAME="auth_service"
export PORT="8081"

# Собрать и запустить сервис
cd auth-service/backend
go run cmd/main.go
```

### Windows (PowerShell)

```powershell
# Запустить PostgreSQL в Docker
docker run --name auth-db -e POSTGRES_USER=decanat_user -e POSTGRES_PASSWORD=decanat_password -e POSTGRES_DB=auth_service -p 5432:5432 -d postgres:latest

# Инициализировать БД
Get-Content auth-service/scripts/init_db.sql | docker exec -i auth-db psql -U decanat_user -d auth_service

# Задать необходимые переменные окружения
$env:JWT_SECRET = "virtual_decanat_secret_key_for_development"
$env:JWT_EXPIRATION = "24h"
$env:DB_HOST = "localhost"
$env:DB_PORT = "5432"
$env:DB_USER = "decanat_user"
$env:DB_PASSWORD = "decanat_password"
$env:DB_NAME = "auth_service"
$env:PORT = "8081"

# Собрать и запустить сервис
cd auth-service/backend
go run cmd/main.go
```

## Способ 3: Быстрый запуск скриптом

Выполните команду:

```bash
# Для Linux/macOS
./demo.sh

# Для Windows
.\demo.ps1
```

## Тестирование API

После запуска сервиса можно тестировать API через curl, Postman или другой инструмент:

### 1. Регистрация пользователя

```bash
curl -X POST http://localhost:8081/register \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser", "email":"test@example.com", "password":"password123", "first_name":"Test", "last_name":"User", "role":"student", "group":"IS-21"}'
```

### 2. Аутентификация

```bash
curl -X POST http://localhost:8081/login \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser", "password":"password123"}'
```

### 3. Получение профиля (требуется токен)

```bash
curl -X GET http://localhost:8081/profile \
     -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### 4. Обновление профиля (требуется токен)

```bash
curl -X PUT http://localhost:8081/profile \
     -H "Authorization: Bearer YOUR_TOKEN_HERE" \
     -H "Content-Type: application/json" \
     -d '{"first_name":"Updated", "last_name":"Name", "faculty":"Updated Faculty"}'
```

### 5. Изменение пароля (требуется токен)

```bash
curl -X POST http://localhost:8081/change-password \
     -H "Authorization: Bearer YOUR_TOKEN_HERE" \
     -H "Content-Type: application/json" \
     -d '{"old_password":"password123", "new_password":"newpassword123"}'
``` 