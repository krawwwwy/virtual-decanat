\c applicant_service;

CREATE TABLE IF NOT EXISTS applications (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    otchestvo VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    number VARCHAR(20) NOT NULL,
    faculty VARCHAR(255) NOT NULL,
    snils VARCHAR(11) NOT NULL,
    passport_seial INT NOT NULL,
    passport_number INT NOT NULL,
    passport_date DATE NOT NULL,
    passport_adress TEXT NOT NULL,
    russian INT NOT NULL CHECK (russian >= 1 AND russian <= 100),
    math INT NOT NULL CHECK (math >= 1 AND math <= 100),
    physics INT CHECK (physics >= 0 AND physics <= 100),
    informatics INT CHECK (informatics >= 0 AND informatics <= 100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for common queries
CREATE INDEX idx_email ON applications(email);
CREATE INDEX idx_snils ON applications(snils);

