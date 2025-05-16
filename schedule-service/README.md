# Schedule Service

Микросервис для управления расписанием, группами и предметами.

## Требования
- Go 1.21+
- Docker и Docker Compose
- PostgreSQL

## Запуск

1. Клонируйте репозиторий:
```bash
git clone https://github.com/krawwwwy/virtual-decanat.git
cd virtual-decanat/schedule-service
```

2. Запустите через Docker Compose:
```bash
docker-compose up -d
```

Сервис будет доступен по адресу: http://localhost:8082

## API Endpoints

### Группы
- POST   /api/v1/group/ — создать
- PUT    /api/v1/group/:id — обновить
- DELETE /api/v1/group/:id — удалить
- GET    /api/v1/group/:id — получить по id
- GET    /api/v1/group/ — список

### Предметы
- POST   /api/v1/subject/ — создать
- PUT    /api/v1/subject/:id — обновить
- DELETE /api/v1/subject/:id — удалить
- GET    /api/v1/subject/:id — получить по id
- GET    /api/v1/subject/ — список

### Расписание
- POST   /api/v1/schedule/ — создать
- PUT    /api/v1/schedule/:id — обновить
- DELETE /api/v1/schedule/:id — удалить
- GET    /api/v1/schedule/:id — получить по id
- GET    /api/v1/schedule/teacher/:teacher_id — по преподавателю
- GET    /api/v1/schedule/group/:group_id — по группе

## Тестирование

Для ручного теста доступен UI: http://localhost:8082/test.html

## Пример запроса на создание пары
```json
{
  "subject_id": 1,
  "teacher_id": 1,
  "group_id": 1,
  "day_of_week": 1,
  "start_time": "2024-05-17T09:00:00Z",
  "end_time": "2024-05-17T10:30:00Z",
  "room": "101"
}
``` 