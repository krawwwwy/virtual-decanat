package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/model"
	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/repository"
)

// JWTConfig конфигурация JWT
type JWTConfig struct {
	Secret     string
	ExpireTime time.Duration
}

// ApplicantService сервис для работы с абитуриентами
type ApplicantService interface {
	Register(ctx context.Context, req *model.RegisterRequest) (*model.Applicant, error)
	Login(ctx context.Context, req *model.LoginRequest) (*model.AuthResponse, error)
	GetApplicant(ctx context.Context, id uint) (*model.Applicant, error)
	CreateApplication(ctx context.Context, applicantID uint, req *model.ApplicationRequest) (*model.Application, error)
	GetApplication(ctx context.Context, id uint) (*model.Application, error)
	GetApplicationsByApplicantID(ctx context.Context, applicantID uint) ([]*model.Application, error)
	UpdateApplication(ctx context.Context, id uint, req *model.ApplicationRequest) (*model.Application, error)
	SubmitApplication(ctx context.Context, id uint) error
	GetApplicationStatus(ctx context.Context, id uint) (model.ApplicationStatus, error)
}

// ApplicantServiceImpl реализация ApplicantService
type ApplicantServiceImpl struct {
	applicantRepo   repository.ApplicantRepository
	applicationRepo repository.ApplicationRepository
	jwtConfig       JWTConfig
}

// NewApplicantService создает новый ApplicantService
func NewApplicantService(applicantRepo repository.ApplicantRepository, applicationRepo repository.ApplicationRepository, jwtConfig JWTConfig) ApplicantService {
	return &ApplicantServiceImpl{
		applicantRepo:   applicantRepo,
		applicationRepo: applicationRepo,
		jwtConfig:       jwtConfig,
	}
}

// Register регистрирует нового абитуриента
func (s *ApplicantServiceImpl) Register(ctx context.Context, req *model.RegisterRequest) (*model.Applicant, error) {
	// Проверяем, что email не занят
	exists, err := s.applicantRepo.Exists(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, model.ErrEmailTaken
	}

	// Создаем нового абитуриента
	applicant, err := model.NewApplicant(req)
	if err != nil {
		return nil, err
	}

	// Сохраняем в базу данных
	if err := s.applicantRepo.Create(ctx, applicant); err != nil {
		return nil, err
	}

	return applicant, nil
}

// Login авторизует абитуриента
func (s *ApplicantServiceImpl) Login(ctx context.Context, req *model.LoginRequest) (*model.AuthResponse, error) {
	// Находим абитуриента по email
	applicant, err := s.applicantRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, model.ErrInvalidCredentials
		}
		return nil, err
	}

	// Проверяем пароль
	if !applicant.CheckPassword(req.Password) {
		return nil, model.ErrInvalidCredentials
	}

	// Создаем JWT токен
	expiresAt := time.Now().Add(s.jwtConfig.ExpireTime)
	token, err := s.createJWTToken(applicant.ID, expiresAt)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		Token:     token,
		ExpiresAt: expiresAt,
		Applicant: applicant,
	}, nil
}

// GetApplicant возвращает абитуриента по ID
func (s *ApplicantServiceImpl) GetApplicant(ctx context.Context, id uint) (*model.Applicant, error) {
	return s.applicantRepo.FindByID(ctx, id)
}

// CreateApplication создает новое заявление
func (s *ApplicantServiceImpl) CreateApplication(ctx context.Context, applicantID uint, req *model.ApplicationRequest) (*model.Application, error) {
	// Проверяем существование абитуриента
	_, err := s.applicantRepo.FindByID(ctx, applicantID)
	if err != nil {
		return nil, err
	}

	// Создаем заявление
	application := &model.Application{
		ApplicantID:        int(applicantID),
		Faculty:            req.Faculty,
		Program:            req.Program,
		Status:             model.StatusDraft,
		DocumentsSubmitted: false,
		PersonalInfo:       req.PersonalInfo,
		EducationInfo:      req.EducationInfo,
		Comments:           req.Comments,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	// Сохраняем в базу данных
	if err := s.applicationRepo.Create(ctx, application); err != nil {
		return nil, err
	}

	return application, nil
}

// GetApplication возвращает заявление по ID
func (s *ApplicantServiceImpl) GetApplication(ctx context.Context, id uint) (*model.Application, error) {
	return s.applicationRepo.FindByID(ctx, id)
}

// GetApplicationsByApplicantID возвращает все заявления абитуриента
func (s *ApplicantServiceImpl) GetApplicationsByApplicantID(ctx context.Context, applicantID uint) ([]*model.Application, error) {
	return s.applicationRepo.FindByApplicantID(ctx, applicantID)
}

// UpdateApplication обновляет заявление
func (s *ApplicantServiceImpl) UpdateApplication(ctx context.Context, id uint, req *model.ApplicationRequest) (*model.Application, error) {
	// Находим заявление
	application, err := s.applicationRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Обновляем только если заявление в статусе черновика
	if application.Status != model.StatusDraft {
		return nil, errors.New("can only update draft applications")
	}

	// Обновляем поля
	application.Faculty = req.Faculty
	application.Program = req.Program
	application.PersonalInfo = req.PersonalInfo
	application.EducationInfo = req.EducationInfo
	application.Comments = req.Comments
	application.UpdatedAt = time.Now()

	// Сохраняем изменения
	if err := s.applicationRepo.Update(ctx, application); err != nil {
		return nil, err
	}

	return application, nil
}

// SubmitApplication отправляет заявление
func (s *ApplicantServiceImpl) SubmitApplication(ctx context.Context, id uint) error {
	return s.applicationRepo.SubmitApplication(ctx, id)
}

// GetApplicationStatus возвращает статус заявления
func (s *ApplicantServiceImpl) GetApplicationStatus(ctx context.Context, id uint) (model.ApplicationStatus, error) {
	return s.applicationRepo.GetApplicationStatus(ctx, id)
}

// createJWTToken создает JWT токен
func (s *ApplicantServiceImpl) createJWTToken(applicantID int, expiresAt time.Time) (string, error) {
	// Создаем заявку на токен с необходимыми полями
	claims := jwt.MapClaims{
		"sub": applicantID,
		"exp": expiresAt.Unix(),
		"iat": time.Now().Unix(),
	}

	// Создаем токен с алгоритмом HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString([]byte(s.jwtConfig.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}