-- Убедитесь, что вы подключены к нужной базе данных
\c auth_lab;

-- Убедимся, что схема общедоступна
CREATE SCHEMA IF NOT EXISTS public;

-- Создание таблиц
CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO roles (name) VALUES ('admin'), ('student'), ('teacher');

-- Создадим таблицу пользователей
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    role VARCHAR(50) NOT NULL,
    group_name VARCHAR(50),
    faculty VARCHAR(100),
    department VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Создадим индексы для повышения производительности
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_role ON users (role);

-- Добавим администратора
INSERT INTO users (username, email, password_hash, first_name, last_name, middle_name, role, created_at, updated_at)
VALUES (
    'admin',
    'admin@example.com',
    -- Хэш пароля "admin123" (bcrypt)
    '$2a$10$RDpDY7qOSVpXOd8uSbHSkOsebPlGq1/dbH2b8VeAmugJXQKgqLMSi',
    'Администратор',
    'Системы',
    NULL,
    'admin',
    NOW(),
    NOW()
)
ON CONFLICT (username) DO NOTHING;

-- Добавим тестовых пользователей (если они не существуют)
-- Студент
INSERT INTO users (username, email, password_hash, first_name, last_name, middle_name, role, group_name, faculty, created_at, updated_at)
VALUES (
    'student',
    'student@example.com',
    -- Хэш пароля "password" (bcrypt)
    '$2a$10$RRYVrFnpYxdeQuJ7fEvfwelcjtgwZmGLfgr7nev5JvMT/5xyyUfhu',
    'Иван',
    'Студентов',
    'Петрович',
    'student',
    'ИС-31',
    'Информационные технологии',
    NOW(),
    NOW()
)
ON CONFLICT (username) DO NOTHING;

-- Преподаватель
INSERT INTO users (username, email, password_hash, first_name, last_name, middle_name, role, department, faculty, created_at, updated_at)
VALUES (
    'teacher',
    'teacher@example.com',
    -- Хэш пароля "password" (bcrypt)
    '$2a$10$RRYVrFnpYxdeQuJ7fEvfwelcjtgwZmGLfgr7nev5JvMT/5xyyUfhu',
    'Мария',
    'Преподавателева',
    'Сергеевна',
    'teacher',
    'Кафедра информационных систем',
    'Информационные технологии',
    NOW(),
    NOW()
)
ON CONFLICT (username) DO NOTHING;

-- Сотрудник деканата
INSERT INTO users (username, email, password_hash, first_name, last_name, middle_name, role, department, faculty, created_at, updated_at)
VALUES (
    'decanat',
    'decanat@example.com',
    -- Хэш пароля "password" (bcrypt)
    '$2a$10$RRYVrFnpYxdeQuJ7fEvfwelcjtgwZmGLfgr7nev5JvMT/5xyyUfhu',
    'Елена',
    'Деканова',
    'Александровна',
    'decanat',
    'Деканат',
    'Информационные технологии',
    NOW(),
    NOW()
)
ON CONFLICT (username) DO NOTHING;

-- Абитуриент
INSERT INTO users (username, email, password_hash, first_name, last_name, middle_name, role, created_at, updated_at)
VALUES (
    'applicant',
    'applicant@example.com',
    -- Хэш пароля "password" (bcrypt)
    '$2a$10$RRYVrFnpYxdeQuJ7fEvfwelcjtgwZmGLfgr7nev5JvMT/5xyyUfhu',
    'Алексей',
    'Абитуриентов',
    'Дмитриевич',
    'applicant',
    NOW(),
    NOW()
)
ON CONFLICT (username) DO NOTHING;

CREATE TABLE IF NOT EXISTS user_roles (
    user_id INT,
    role_id INT,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);
