# Установим переменные окружения для демонстрации
$env:DB_USER = "decanat_user"
$env:DB_PASSWORD = "decanat_password"
$env:DB_NAME = "auth_service" 
$env:DB_HOST = "localhost"
$env:DB_PORT = "5432"
$env:JWT_SECRET = "virtual_decanat_secret_key_for_development"
$env:JWT_EXPIRATION = "24h"
$env:PORT = "8081"

# Создаем базу данных через Docker
Write-Host "Запуск PostgreSQL для демонстрации..."
docker run --name auth-db -e POSTGRES_USER=$env:DB_USER -e POSTGRES_PASSWORD=$env:DB_PASSWORD -e POSTGRES_DB=$env:DB_NAME -p 5432:5432 -d postgres:latest

# Даем PostgreSQL время на инициализацию
Start-Sleep -Seconds 5

# Инициализируем БД
Write-Host "Инициализация базы данных..."
Get-Content auth-service/scripts/init_db.sql | docker exec -i auth-db psql -U $env:DB_USER -d $env:DB_NAME

Write-Host "Сборка сервиса аутентификации..."
Set-Location -Path auth-service/backend
go build -o auth-service.exe cmd/main.go

Write-Host "Запуск сервиса аутентификации на порту $env:PORT..."
$authProcess = Start-Process -FilePath ".\auth-service.exe" -PassThru -NoNewWindow

# Даем сервису время на запуск
Start-Sleep -Seconds 2

Write-Host "Демонстрация API endpoints:"

Write-Host "`n1. Регистрация нового пользователя:"
$registerResponse = Invoke-RestMethod -Uri "http://localhost:$env:PORT/register" -Method Post -Headers @{"Content-Type" = "application/json"} -Body '{"username":"testuser", "email":"test@example.com", "password":"password123", "first_name":"Test", "last_name":"User", "role":"student", "group":"ИС-21"}' -ErrorAction SilentlyContinue
$registerResponse | ConvertTo-Json

Write-Host "`n2. Аутентификация пользователя:"
$loginResponse = Invoke-RestMethod -Uri "http://localhost:$env:PORT/login" -Method Post -Headers @{"Content-Type" = "application/json"} -Body '{"username":"testuser", "password":"password123"}' -ErrorAction SilentlyContinue
$token = $loginResponse.token
Write-Host "Получен токен: $token"

Write-Host "`n3. Получение профиля пользователя:"
$profileResponse = Invoke-RestMethod -Uri "http://localhost:$env:PORT/profile" -Method Get -Headers @{"Authorization" = "Bearer $token"} -ErrorAction SilentlyContinue
$profileResponse | ConvertTo-Json

Write-Host "`n4. Обновление профиля пользователя:"
$updateResponse = Invoke-RestMethod -Uri "http://localhost:$env:PORT/profile" -Method Put -Headers @{"Authorization" = "Bearer $token"; "Content-Type" = "application/json"} -Body '{"first_name":"Updated", "last_name":"Name", "faculty":"Обновленный факультет"}' -ErrorAction SilentlyContinue
$updateResponse | ConvertTo-Json

Write-Host "`n5. Изменение пароля пользователя:"
$passwordResponse = Invoke-RestMethod -Uri "http://localhost:$env:PORT/change-password" -Method Post -Headers @{"Authorization" = "Bearer $token"; "Content-Type" = "application/json"} -Body '{"old_password":"password123", "new_password":"newpassword123"}' -ErrorAction SilentlyContinue
$passwordResponse | ConvertTo-Json

# Останавливаем сервис
Write-Host "`nОстановка сервиса аутентификации..."
Stop-Process -Id $authProcess.Id

# Возвращаемся в корневую директорию
Set-Location -Path ../..

# Останавливаем и удаляем Docker-контейнер
Write-Host "Остановка и удаление контейнера с базой данных..."
docker stop auth-db
docker rm auth-db

Write-Host "`nДемонстрация завершена!" 