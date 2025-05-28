package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User представляет пользователя системы
type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Email        string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	FirstName    string    `json:"first_name" gorm:"not null"`
	LastName     string    `json:"last_name" gorm:"not null"`
	MiddleName   string    `json:"middle_name"`
	RoleID       uint      `json:"role_id" gorm:"not null"`
	Role         *Role     `json:"role,omitempty" gorm:"foreignKey:RoleID"`
	IsActive     bool      `json:"is_active" gorm:"default:true"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели User
func (User) TableName() string {
	return "auth.users"
}

// SetPassword устанавливает хэш пароля
func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

// CheckPassword проверяет соответствие пароля хэшу
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

// Role представляет роль пользователя
type Role struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Role
func (Role) TableName() string {
	return "auth.roles"
}

// Student представляет студента
type Student struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	User      *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	GroupID   uint      `json:"group_id" gorm:"not null"`
	StudentID string    `json:"student_id" gorm:"unique;not null"`
	BirthDate time.Time `json:"birth_date"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Student
func (Student) TableName() string {
	return "auth.students"
}

// Teacher представляет преподавателя
type Teacher struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	User       *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Department string    `json:"department" gorm:"not null"`
	Position   string    `json:"position" gorm:"not null"`
	Degree     string    `json:"degree"`
	BirthDate  time.Time `json:"birth_date"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Teacher
func (Teacher) TableName() string {
	return "auth.teachers"
}

// Staff представляет сотрудника деканата
type Staff struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	UserID        uint      `json:"user_id" gorm:"not null"`
	User          *User     `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Department    string    `json:"department" gorm:"not null"`
	Position      string    `json:"position" gorm:"not null"`
	InternalPhone string    `json:"internal_phone"`
	Gender        string    `json:"gender"`
	BirthDate     time.Time `json:"birth_date"`
	Phone         string    `json:"phone"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Staff
func (Staff) TableName() string {
	return "auth.staff"
} 