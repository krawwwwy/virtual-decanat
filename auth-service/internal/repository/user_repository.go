package repository

import (
	"context"
	"errors"

	"github.com/krawwwwy/virtual-decanat/auth-service/internal/model"
	"gorm.io/gorm"
)

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id uint) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint) error
	FindRoleByName(ctx context.Context, name string) (*model.Role, error)
	CreateStudent(ctx context.Context, student *model.Student) error
	CreateTeacher(ctx context.Context, teacher *model.Teacher) error
	FindStudentByUserID(ctx context.Context, userID uint) (*model.Student, error)
	FindTeacherByUserID(ctx context.Context, userID uint) (*model.Teacher, error)
}

// UserRepositoryImpl реализация UserRepository
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository создает новый UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// Create создает нового пользователя
func (r *UserRepositoryImpl) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// FindByID находит пользователя по ID
func (r *UserRepositoryImpl) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Role").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// FindByEmail находит пользователя по email
func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Role").Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Update обновляет пользователя
func (r *UserRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete удаляет пользователя
func (r *UserRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
}

// FindRoleByName находит роль по имени
func (r *UserRepositoryImpl) FindRoleByName(ctx context.Context, name string) (*model.Role, error) {
	var role model.Role
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

// CreateStudent создает нового студента
func (r *UserRepositoryImpl) CreateStudent(ctx context.Context, student *model.Student) error {
	return r.db.WithContext(ctx).Create(student).Error
}

// CreateTeacher создает нового преподавателя
func (r *UserRepositoryImpl) CreateTeacher(ctx context.Context, teacher *model.Teacher) error {
	return r.db.WithContext(ctx).Create(teacher).Error
}

// FindStudentByUserID находит студента по ID пользователя
func (r *UserRepositoryImpl) FindStudentByUserID(ctx context.Context, userID uint) (*model.Student, error) {
	var student model.Student
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&student).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &student, nil
}

// FindTeacherByUserID находит преподавателя по ID пользователя
func (r *UserRepositoryImpl) FindTeacherByUserID(ctx context.Context, userID uint) (*model.Teacher, error) {
	var teacher model.Teacher
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&teacher).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &teacher, nil
} 