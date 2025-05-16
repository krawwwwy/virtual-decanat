# Auth Service

Сервис аутентификации и авторизации для системы "Виртуальный деканат".

## Требования

- Go 1.21+
- Docker и Docker Compose
- PostgreSQL

## Запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/krawwwwy/virtual-decanat.git
cd virtual-decanat/auth-service
```

2. Запустите через Docker Compose:
```bash
docker-compose up -d
```

Сервис будет доступен по адресу: http://localhost:8081

## API Endpoints

### Регистрация
```http
POST /api/v1/auth/register
Content-Type: application/json

{
    "email": "test@example.com",
    "password": "Test123!@#",
    "first_name": "Test",
    "last_name": "User",
    "role": "student"
}
```

### Вход
```http
POST /api/v1/auth/login
Content-Type: application/json

{
    "email": "test@example.com",
    "password": "Test123!@#"
}
```

### Обновление токена
```http
POST /api/v1/auth/refresh
Content-Type: application/json

{
    "refresh_token": "your_refresh_token"
}
```

### Получение профиля
```http
GET /api/v1/users/profile
Authorization: Bearer your_access_token
```

### Обновление профиля
```http
PUT /api/v1/users/profile
Authorization: Bearer your_access_token
Content-Type: application/json

{
    "first_name": "New Name",
    "last_name": "New Last Name"
}
```

### Смена пароля
```http
POST /api/v1/users/change-password
Authorization: Bearer your_access_token
Content-Type: application/json

{
    "old_password": "old_password",
    "new_password": "new_password"
}
```

## Тестирование

Для тестирования API доступен веб-интерфейс: http://localhost:8081/test.html

## Требования к паролям

- Минимум 8 символов для регистрации
- Минимум 6 символов для входа
- Email должен быть валидным (содержать @)

## Роли пользователей

- student - Студент
- teacher - Преподаватель
- dean_office - Деканат
- admin - Администратор
- applicant - Абитуриент 