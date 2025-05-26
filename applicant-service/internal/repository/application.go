package repository

import (
	"context"

	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/model"
	"gorm.io/gorm"
)

// ApplicationRepository интерфейс для работы с заявлениями
type ApplicationRepository interface {
	Create(ctx context.Context, application *model.Application) error
	FindByID(ctx context.Context, id uint) (*model.Application, error)
	FindByApplicantID(ctx context.Context, applicantID uint) ([]*model.Application, error)
	Update(ctx context.Context, application *model.Application) error
	Delete(ctx context.Context, id uint) error
	GetApplicationStatus(ctx context.Context, id uint) (model.ApplicationStatus, error)
	SubmitApplication(ctx context.Context, id uint) error
	UpdateStatus(ctx context.Context, id uint, status model.ApplicationStatus) error
}

// ApplicationRepositoryImpl реализация ApplicationRepository
type ApplicationRepositoryImpl struct {
	db *gorm.DB
}

// NewApplicationRepository создает новый ApplicationRepository
func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &ApplicationRepositoryImpl{db: db}
}

// Create создает новое заявление
func (r *ApplicationRepositoryImpl) Create(ctx context.Context, application *model.Application) error {
	// Устанавливаем статус "черновик" по умолчанию
	if application.Status == "" {
		application.Status = model.StatusDraft
	}
	
	return r.db.WithContext(ctx).Create(application).Error
}

// Update обновляет заявление
func (r *ApplicationRepositoryImpl) Update(ctx context.Context, application *model.Application) error {
	return r.db.WithContext(ctx).Save(application).Error
}

// Delete удаляет заявление
func (r *ApplicationRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Application{}, id).Error
}

// FindByID находит заявление по ID
func (r *ApplicationRepositoryImpl) FindByID(ctx context.Context, id uint) (*model.Application, error) {
	var application model.Application
	if err := r.db.WithContext(ctx).First(&application, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, model.ErrNotFound
		}
		return nil, err
	}
	return &application, nil
}

// FindByApplicantID находит все заявления абитуриента по его ID
func (r *ApplicationRepositoryImpl) FindByApplicantID(ctx context.Context, applicantID uint) ([]*model.Application, error) {
	var applications []*model.Application
	if err := r.db.WithContext(ctx).Where("applicant_id = ?", applicantID).Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

// GetApplicationStatus возвращает статус заявления
func (r *ApplicationRepositoryImpl) GetApplicationStatus(ctx context.Context, id uint) (model.ApplicationStatus, error) {
	var application model.Application
	if err := r.db.WithContext(ctx).Select("status").First(&application, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", model.ErrNotFound
		}
		return "", err
	}
	return application.Status, nil
}

// SubmitApplication отправляет заявление (меняет статус с черновика на отправлено)
func (r *ApplicationRepositoryImpl) SubmitApplication(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Model(&model.Application{}).
		Where("id = ? AND status = ?", id, model.StatusDraft).
		Update("status", model.StatusSubmitted)
	
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return model.ErrNotFound
	}
	
	return nil
}

// UpdateStatus обновляет статус заявления
func (r *ApplicationRepositoryImpl) UpdateStatus(ctx context.Context, id uint, status model.ApplicationStatus) error {
	result := r.db.WithContext(ctx).Model(&model.Application{}).
		Where("id = ?", id).
		Update("status", status)
	
	if result.Error != nil {
		return result.Error
	}
	
	if result.RowsAffected == 0 {
		return model.ErrNotFound
	}
	
	return nil
}

