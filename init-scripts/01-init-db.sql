-- Создание схем для каждого микросервиса
CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS schedule;
CREATE SCHEMA IF NOT EXISTS club;
CREATE SCHEMA IF NOT EXISTS performance;
CREATE SCHEMA IF NOT EXISTS applicant;
CREATE SCHEMA IF NOT EXISTS support;

-- Создание таблицы ролей
CREATE TABLE IF NOT EXISTS auth.roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы пользователей
CREATE TABLE IF NOT EXISTS auth.users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    role_id INTEGER REFERENCES auth.roles(id),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы групп
CREATE TABLE IF NOT EXISTS schedule.groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    faculty VARCHAR(100) NOT NULL,
    year INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы студентов
CREATE TABLE IF NOT EXISTS auth.students (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES auth.users(id),
    group_id INTEGER REFERENCES schedule.groups(id),
    student_id VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы преподавателей
CREATE TABLE IF NOT EXISTS auth.teachers (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES auth.users(id),
    department VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы дисциплин
CREATE TABLE IF NOT EXISTS schedule.subjects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(20) NOT NULL UNIQUE,
    credits INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы расписания
CREATE TABLE IF NOT EXISTS schedule.schedule (
    id SERIAL PRIMARY KEY,
    subject_id INTEGER REFERENCES schedule.subjects(id),
    teacher_id INTEGER REFERENCES auth.teachers(id),
    group_id INTEGER REFERENCES schedule.groups(id),
    day_of_week INTEGER NOT NULL,
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    room VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы оценок
CREATE TABLE IF NOT EXISTS performance.grades (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES auth.students(id),
    subject_id INTEGER REFERENCES schedule.subjects(id),
    teacher_id INTEGER REFERENCES auth.teachers(id),
    grade INTEGER NOT NULL,
    semester INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы посещаемости
CREATE TABLE IF NOT EXISTS performance.attendance (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES auth.students(id),
    schedule_id INTEGER REFERENCES schedule.schedule(id),
    date DATE NOT NULL,
    is_present BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы задолженностей
CREATE TABLE IF NOT EXISTS performance.debts (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES auth.students(id),
    subject_id INTEGER REFERENCES schedule.subjects(id),
    description TEXT NOT NULL,
    deadline DATE,
    is_resolved BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы студенческих клубов
CREATE TABLE IF NOT EXISTS club.clubs (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Удаляем старую таблицу memberships, если есть
DROP TABLE IF EXISTS club.memberships;

-- Создание таблицы членства в клубах (актуальная структура)
CREATE TABLE IF NOT EXISTS club.club_members (
    id SERIAL PRIMARY KEY,
    club_id INTEGER REFERENCES club.clubs(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES auth.users(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL,
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы заявок на вступление в клуб
CREATE TABLE IF NOT EXISTS club.applications (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES auth.students(id),
    club_id INTEGER REFERENCES club.clubs(id),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    message TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы абитуриентов
CREATE TABLE IF NOT EXISTS applicant.applicants (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    email VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы заявлений абитуриентов
CREATE TABLE IF NOT EXISTS applicant.applications (
    id SERIAL PRIMARY KEY,
    applicant_id INTEGER REFERENCES applicant.applicants(id),
    faculty VARCHAR(100) NOT NULL,
    program VARCHAR(255) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    documents_submitted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы типов социальной поддержки
CREATE TABLE IF NOT EXISTS support.support_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Создание таблицы заявлений на социальную поддержку
CREATE TABLE IF NOT EXISTS support.applications (
    id SERIAL PRIMARY KEY,
    student_id INTEGER REFERENCES auth.students(id),
    support_type_id INTEGER REFERENCES support.support_types(id),
    status VARCHAR(20) NOT NULL DEFAULT 'pending',
    documents_submitted BOOLEAN DEFAULT FALSE,
    comment TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Вставка начальных данных для ролей
INSERT INTO auth.roles (name) VALUES 
    ('admin'),
    ('student'),
    ('teacher'),
    ('dean_office'),
    ('applicant')
ON CONFLICT (name) DO NOTHING;

-- Вставка тестовой группы для студентов
INSERT INTO schedule.groups (id, name, faculty, year) VALUES 
    (1, 'Test Group', 'Test Faculty', 1); 