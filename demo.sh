#!/bin/bash

# Установим переменные окружения для демонстрации
export DB_USER=decanat_user
export DB_PASSWORD=decanat_password
export DB_NAME=auth_service
export DB_HOST=localhost
export DB_PORT=5432
export JWT_SECRET=virtual_decanat_secret_key_for_development
export JWT_EXPIRATION=24h
export PORT=8081

# Создаем базу данных через Docker
echo "Запуск PostgreSQL для демонстрации..."
docker run --name auth-db -e POSTGRES_USER=$DB_USER -e POSTGRES_PASSWORD=$DB_PASSWORD -e POSTGRES_DB=$DB_NAME -p 5432:5432 -d postgres:latest

# Даем PostgreSQL время на инициализацию
sleep 5

# Инициализируем БД
echo "Инициализация базы данных..."
docker exec -i auth-db psql -U $DB_USER -d $DB_NAME < auth-service/scripts/init_db.sql

echo "Сборка сервиса аутентификации..."
cd auth-service/backend
go build -o auth-service cmd/main.go

echo "Запуск сервиса аутентификации на порту $PORT..."
./auth-service &
AUTH_PID=$!

# Даем сервису время на запуск
sleep 2

echo "Демонстрация API endpoints:"

echo -e "\n1. Регистрация нового пользователя:"
curl -X POST http://localhost:$PORT/register \
     -H "Content-Type: application/json" \
     -d '{"username":"testuser", "email":"test@example.com", "password":"password123", "first_name":"Test", "last_name":"User", "role":"student", "group":"ИС-21"}'

echo -e "\n\n2. Аутентификация пользователя:"
TOKEN=$(curl -s -X POST http://localhost:$PORT/login \
        -H "Content-Type: application/json" \
        -d '{"username":"testuser", "password":"password123"}' | grep -o '"token":"[^"]*' | cut -d'"' -f4)

echo "Получен токен: $TOKEN"

echo -e "\n3. Получение профиля пользователя:"
curl -X GET http://localhost:$PORT/profile \
     -H "Authorization: Bearer $TOKEN"

echo -e "\n\n4. Обновление профиля пользователя:"
curl -X PUT http://localhost:$PORT/profile \
     -H "Authorization: Bearer $TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"first_name":"Updated", "last_name":"Name", "faculty":"Обновленный факультет"}'

echo -e "\n\n5. Изменение пароля пользователя:"
curl -X POST http://localhost:$PORT/change-password \
     -H "Authorization: Bearer $TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"old_password":"password123", "new_password":"newpassword123"}'

# Останавливаем сервис
echo -e "\n\nОстановка сервиса аутентификации..."
kill $AUTH_PID

# Останавливаем и удаляем Docker-контейнер
echo "Остановка и удаление контейнера с базой данных..."
docker stop auth-db
docker rm auth-db

echo -e "\nДемонстрация завершена!" 