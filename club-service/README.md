# Сервис Клубов

Этот микросервис управляет студенческими клубами и их участниками в системе Виртуальный Деканат.

## Возможности

- Создание, просмотр, обновление и удаление клубов
- Добавление и удаление участников клубов
- Просмотр участников клубов
- Управление ролями участников (администратор/участник)

## API Endpoints

### Клубы

- `POST /api/clubs` - Создать новый клуб
- `GET /api/clubs` - Получить все клубы
- `GET /api/clubs/:id` - Получить клуб по ID
- `PUT /api/clubs/:id` - Обновить клуб
- `DELETE /api/clubs/:id` - Удалить клуб

### Участники

- `POST /api/members` - Добавить участника в клуб
- `GET /api/members/club/:clubId` - Получить всех участников клуба
- `DELETE /api/members/:id` - Удалить участника из клуба

## Схема Базы Данных

### Таблица Клубов
```sql
CREATE TABLE clubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Таблица Участников Клуба
```sql
CREATE TABLE club_members (
    id SERIAL PRIMARY KEY,
    club_id INTEGER REFERENCES clubs(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL,
    role VARCHAR(50) NOT NULL,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Переменные Окружения

- `DB_HOST` - Хост базы данных (по умолчанию: localhost)
- `DB_PORT` - Порт базы данных (по умолчанию: 5432)
- `DB_USER` - Пользователь базы данных
- `DB_PASSWORD` - Пароль базы данных
- `DB_NAME` - Имя базы данных
- `PORT` - Порт сервиса (по умолчанию: 8083)

## Запуск Сервиса

1. Собрать Docker образ:
```bash
docker build -t club-service .
```

2. Запустить контейнер:
```bash
docker run -p 8083:8083 \
  -e DB_HOST=postgres \
  -e DB_PORT=5432 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=postgres \
  -e DB_NAME=decanat \
  club-service
```

## Тестирование

Сервис включает веб-интерфейс для тестирования по адресу `http://localhost:8083`. Вы можете использовать этот интерфейс для:

1. Создания новых клубов
2. Просмотра всех клубов
3. Добавления участников в клубы
4. Просмотра участников клубов
5. Удаления участников из клубов
6. Удаления клубов 