-- Create auth schema
CREATE SCHEMA IF NOT EXISTS auth;

-- Create roles table
CREATE TABLE IF NOT EXISTS auth.roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create users table
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

-- Create schedule schema
CREATE SCHEMA IF NOT EXISTS schedule;

-- Create groups table
CREATE TABLE IF NOT EXISTS schedule.groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    faculty VARCHAR(100) NOT NULL,
    year INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create students table
CREATE TABLE IF NOT EXISTS auth.students (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES auth.users(id),
    group_id INTEGER REFERENCES schedule.groups(id),
    student_id VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create teachers table
CREATE TABLE IF NOT EXISTS auth.teachers (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES auth.users(id),
    department VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Insert default roles
INSERT INTO auth.roles (name) VALUES 
    ('admin'),
    ('student'),
    ('teacher'),
    ('dean_office'),
    ('applicant')
ON CONFLICT (name) DO NOTHING;

-- Insert a default group for testing
INSERT INTO schedule.groups (name, faculty, year)
VALUES ('Test Group', 'Test Faculty', 2023)
ON CONFLICT (name) DO NOTHING; 