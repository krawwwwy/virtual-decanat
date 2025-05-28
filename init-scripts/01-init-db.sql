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
    birth_date DATE,
    phone VARCHAR(20),
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

-- Создание таблицы сотрудников деканата
CREATE TABLE IF NOT EXISTS auth.staff (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES auth.users(id),
    department VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    internal_phone VARCHAR(20),
    gender VARCHAR(10),
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
DROP TABLE IF EXISTS applicant.applications;
DROP TABLE IF EXISTS applicant.applicants;

CREATE TABLE IF NOT EXISTS applicant.applicants (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    middle_name VARCHAR(100),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
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
    status VARCHAR(50) NOT NULL DEFAULT 'draft',
    documents_submitted BOOLEAN DEFAULT FALSE,
    comments TEXT,
    -- Персональные данные
    passport_series VARCHAR(10),
    passport_number VARCHAR(20),
    passport_issued_by TEXT,
    passport_date TIMESTAMP WITH TIME ZONE,
    birth_date TIMESTAMP WITH TIME ZONE,
    birth_place VARCHAR(255),
    address TEXT,
    -- Данные об образовании
    education_type VARCHAR(100),
    institution VARCHAR(255),
    graduation_year INTEGER,
    document_number VARCHAR(50),
    document_date VARCHAR(50),
    average_grade REAL,
    has_original_documents BOOLEAN DEFAULT FALSE,
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

-- Вставка тестового пользователя
INSERT INTO auth.users (id, email, password_hash, first_name, last_name, role_id) VALUES (1, 'test@student.com', 'testhash', 'Test', 'Student', 2) ON CONFLICT (id) DO NOTHING;
-- Вставка тестового преподавателя
INSERT INTO auth.users (id, email, password_hash, first_name, last_name, role_id) VALUES (2, 'test@teacher.com', 'testhash', 'Test', 'Teacher', 3) ON CONFLICT (id) DO NOTHING;
INSERT INTO auth.teachers (id, user_id, department, position) VALUES (1, 2, 'Test Dept', 'Professor') ON CONFLICT (id) DO NOTHING;
-- Вставка тестового студента
INSERT INTO auth.students (id, user_id, group_id, student_id) VALUES (1, 1, 1, 'S1') ON CONFLICT (id) DO NOTHING;
-- Вставка тестовой дисциплины
INSERT INTO schedule.subjects (id, name, code, credits) VALUES (1, 'Test Subject', 'TS1', 5) ON CONFLICT (id) DO NOTHING;
-- Вставка тестового расписания
INSERT INTO schedule.schedule (id, subject_id, teacher_id, group_id, day_of_week, start_time, end_time, room) VALUES (1, 1, 1, 1, 1, '09:00', '10:30', '101') ON CONFLICT (id) DO NOTHING;

-- Тестовые данные для performance.grades
INSERT INTO performance.grades (student_id, subject_id, teacher_id, grade, semester)
VALUES (1, 1, 1, 85, 1), (1, 1, 1, 90, 2)
ON CONFLICT DO NOTHING;

-- Тестовые данные для performance.attendance
INSERT INTO performance.attendance (student_id, schedule_id, date, is_present)
VALUES (1, 1, '2024-04-01', true), (1, 1, '2024-04-02', false)
ON CONFLICT DO NOTHING;

-- Тестовые данные для performance.debts
INSERT INTO performance.debts (student_id, subject_id, description, deadline, is_resolved)
VALUES (1, 1, 'Не сдал курсовую', '2024-06-01', false)
ON CONFLICT DO NOTHING; 

-- Вставка тестового абитуриента
INSERT INTO applicant.applicants (id, first_name, last_name, middle_name, email, password_hash, phone)
VALUES (1, 'Test', 'Abiturient', 'Testovich', 'test@abiturient.com', 'testhash', '1234567890')
ON CONFLICT DO NOTHING;

-- Вставка тестовой заявки абитуриента
INSERT INTO applicant.applications (
    id, applicant_id, faculty, program, status, documents_submitted,
    passport_series, passport_number, passport_issued_by, passport_date,
    birth_date, birth_place, address,
    education_type, institution, graduation_year, document_number,
    document_date, average_grade, has_original_documents
)
VALUES (
    1, 1, 'Факультет информатики', 'Программная инженерия', 'submitted', true,
    '1234', '567890', 'МВД России', '2015-01-15',
    '2000-05-20', 'Москва', 'ул. Примерная, д. 1, кв. 1',
    'Среднее общее', 'Школа №123', 2022, 'A-123456',
    '15.06.2022', 4.5, true
)
ON CONFLICT DO NOTHING;




