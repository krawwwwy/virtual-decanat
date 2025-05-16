package models

import "time"

// Роли пользователей
const (
	RoleStudent     = "student"     // Студент
	RoleTeacher     = "teacher"     // Преподаватель
	RoleDecanat     = "decanat"     // Сотрудник деканата
	RoleApplicant   = "applicant"   // Абитуриент
	RoleAdmin       = "admin"       // Администратор
)

// User представляет информацию о пользователе в системе
type User struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Не отправляем хэш пароля клиенту
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	MiddleName   string    `json:"middle_name,omitempty"`
	Role         string    `json:"role"`
	Group        string    `json:"group,omitempty"` // Для студентов
	Faculty      string    `json:"faculty,omitempty"`
	Department   string    `json:"department,omitempty"` // Для преподавателей
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// UserRegistration представляет данные для регистрации нового пользователя
type UserRegistration struct {
	Username   string `json:"username" validate:"required,min=3,max=50"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
	MiddleName string `json:"middle_name,omitempty"`
	Role       string `json:"role" validate:"required,oneof=student teacher decanat applicant"`
	Group      string `json:"group,omitempty"`
	Faculty    string `json:"faculty,omitempty"`
	Department string `json:"department,omitempty"`
}

// UserLogin представляет данные для аутентификации пользователя
type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// UserResponse представляет данные о пользователе, возвращаемые клиенту
// без чувствительной информации
type UserResponse struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	MiddleName string    `json:"middle_name,omitempty"`
	Role       string    `json:"role"`
	Group      string    `json:"group,omitempty"`
	Faculty    string    `json:"faculty,omitempty"`
	Department string    `json:"department,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

// TokenResponse представляет ответ с JWT-токеном
type TokenResponse struct {
	Token     string       `json:"token"`
	ExpiresAt time.Time    `json:"expires_at"`
	User      UserResponse `json:"user"`
} 