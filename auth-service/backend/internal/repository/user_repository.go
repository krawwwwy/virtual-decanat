package repository

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/krawwwwy/virtual-decanat/auth-service/internal/models"
)

// Определение ошибок
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

// UserRepository предоставляет методы для работы с пользователями в базе данных
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository создает новый репозиторий пользователей
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser создает нового пользователя в базе данных
func (r *UserRepository) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (username, email, password_hash, first_name, last_name, middle_name, role, 
		                  group_name, faculty, department, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	return r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.MiddleName,
		user.Role,
		user.Group,
		user.Faculty,
		user.Department,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
}

// GetUserByID получает пользователя по ID
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, first_name, last_name, middle_name, role, 
		       group_name, faculty, department, created_at, updated_at
		FROM users 
		WHERE id = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Role,
		&user.Group,
		&user.Faculty,
		&user.Department,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		log.Printf("Failed to get user by ID: %v", err)
		return nil, err
	}

	return user, nil
}

// GetUserByUsername получает пользователя по имени пользователя
func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, first_name, last_name, middle_name, role, 
		       group_name, faculty, department, created_at, updated_at
		FROM users 
		WHERE username = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Role,
		&user.Group,
		&user.Faculty,
		&user.Department,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		log.Printf("Failed to get user by username: %v", err)
		return nil, err
	}

	return user, nil
}

// GetUserByEmail получает пользователя по email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	query := `
		SELECT id, username, email, password_hash, first_name, last_name, middle_name, role, 
		       group_name, faculty, department, created_at, updated_at
		FROM users 
		WHERE email = $1`

	user := &models.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Role,
		&user.Group,
		&user.Faculty,
		&user.Department,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		log.Printf("Failed to get user by email: %v", err)
		return nil, err
	}

	return user, nil
}

// UpdateUser обновляет данные пользователя
func (r *UserRepository) UpdateUser(user *models.User) error {
	query := `
		UPDATE users
		SET email = $1, first_name = $2, last_name = $3, middle_name = $4,
		    role = $5, group_name = $6, faculty = $7, department = $8, updated_at = $9
		WHERE id = $10`

	user.UpdatedAt = time.Now()

	_, err := r.db.Exec(
		query,
		user.Email,
		user.FirstName,
		user.LastName,
		user.MiddleName,
		user.Role,
		user.Group,
		user.Faculty,
		user.Department,
		user.UpdatedAt,
		user.ID,
	)

	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return err
	}

	return nil
}

// UpdateUserPassword обновляет пароль пользователя
func (r *UserRepository) UpdateUserPassword(userID int, passwordHash string) error {
	query := `
		UPDATE users
		SET password_hash = $1, updated_at = $2
		WHERE id = $3`

	_, err := r.db.Exec(query, passwordHash, time.Now(), userID)

	if err != nil {
		log.Printf("Failed to update user password: %v", err)
		return err
	}

	return nil
}

// DeleteUser удаляет пользователя по ID
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Failed to get rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return ErrUserNotFound
	}

	return nil
}

// CheckUsernameExists проверяет, существует ли пользователь с указанным именем
func (r *UserRepository) CheckUsernameExists(username string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE username = $1`

	err := r.db.QueryRow(query, username).Scan(&count)
	if err != nil {
		log.Printf("Failed to check username existence: %v", err)
		return false, err
	}

	return count > 0, nil
}

// CheckEmailExists проверяет, существует ли пользователь с указанным email
func (r *UserRepository) CheckEmailExists(email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM users WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		log.Printf("Failed to check email existence: %v", err)
		return false, err
	}

	return count > 0, nil
} 