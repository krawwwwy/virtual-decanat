package model

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrEmailTaken       = errors.New("email already taken")
	ErrInvalidCredentials = errors.New("invalid email or password")
)

// Applicant представляет абитуриента
type Applicant struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  string    `json:"middle_name"`
	Email       string    `json:"email"`
	PasswordHash string   `json:"-"` // не возвращаем хеш пароля в JSON
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Applicant
func (a *Applicant) TableName() string {
	return "applicant.applicants"
}

// SetPassword устанавливает хэш пароля
func (a *Applicant) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.PasswordHash = string(hash)
	return nil
}

// CheckPassword проверяет соответствие пароля хэшу
func (a *Applicant) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(password))
	return err == nil
}

// NewApplicant создает нового абитуриента из данных запроса
func NewApplicant(req *RegisterRequest) (*Applicant, error) {
	applicant := &Applicant{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		MiddleName: req.MiddleName,
		Email:      req.Email,
		Phone:      req.Phone,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	
	if err := applicant.SetPassword(req.Password); err != nil {
		return nil, err
	}
	
	return applicant, nil
}

// RegisterRequest представляет запрос на регистрацию
type RegisterRequest struct {
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=8"`
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone" binding:"required"`
}

// LoginRequest представляет запрос на вход в систему
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// AuthResponse представляет ответ с токеном авторизации
type AuthResponse struct {
	Token     string     `json:"token"`
	ExpiresAt time.Time  `json:"expires_at"`
	Applicant *Applicant `json:"applicant"`
}





