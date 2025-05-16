-- Insert roles needed for auth service
INSERT INTO auth.roles (name) VALUES 
    ('admin'),
    ('student'),
    ('teacher'),
    ('dean_office'),
    ('applicant')
ON CONFLICT (name) DO NOTHING; 