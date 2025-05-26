package repository

import (
	"context"

	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/model"
	"gorm.io/gorm"
)

var ErrNotFound = model.ErrNotFound

// ApplicantRepository интерфейс для работы с абитуриентами
type ApplicantRepository interface {
	Create(ctx context.Context, applicant *model.Applicant) error
	FindByID(ctx context.Context, id uint) (*model.Applicant, error)
	FindByEmail(ctx context.Context, email string) (*model.Applicant, error)
	FindByApplicationID(ctx context.Context, applicationID uint) (*model.Applicant, error)
	Update(ctx context.Context, applicant *model.Applicant) error
	Delete(ctx context.Context, id uint) error
	Exists(ctx context.Context, email string) (bool, error)
}

// ApplicantRepositoryImpl реализация ApplicantRepository	
type ApplicantRepositoryImpl struct {
	db *gorm.DB
}

// NewApplicantRepository создает новый ApplicantRepository
func NewApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &ApplicantRepositoryImpl{db: db}
}

// Create создает нового абитуриента
func (r *ApplicantRepositoryImpl) Create(ctx context.Context, applicant *model.Applicant) error {
	return r.db.WithContext(ctx).Create(applicant).Error
}

// FindByID находит абитуриента по ID
func (r *ApplicantRepositoryImpl) FindByID(ctx context.Context, id uint) (*model.Applicant, error) {
	var applicant model.Applicant
	if err := r.db.WithContext(ctx).First(&applicant, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return &applicant, nil
}

// Update обновляет данные абитуриента
func (r *ApplicantRepositoryImpl) Update(ctx context.Context, applicant *model.Applicant) error {
	return r.db.WithContext(ctx).Save(applicant).Error
}

// Delete удаляет абитуриента
func (r *ApplicantRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Applicant{}, id).Error
}

// FindByApplicationID находит абитуриента по ID заявления
func (r *ApplicantRepositoryImpl) FindByApplicationID(ctx context.Context, applicationID uint) (*model.Applicant, error) {
	var application model.Application
	if err := r.db.WithContext(ctx).Select("applicant_id").First(&application, applicationID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	
	return r.FindByID(ctx, uint(application.ApplicantID))
}

// FindByEmail находит абитуриента по email
func (r *ApplicantRepositoryImpl) FindByEmail(ctx context.Context, email string) (*model.Applicant, error) {
	var applicant model.Applicant
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&applicant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return &applicant, nil
}

// Exists проверяет существование абитуриента с указанным email
func (r *ApplicantRepositoryImpl) Exists(ctx context.Context, email string) (bool, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.Applicant{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

