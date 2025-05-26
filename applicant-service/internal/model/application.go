package model

import "time"

// ApplicationStatus представляет статус заявления
type ApplicationStatus string

const (
	// StatusDraft заявление в черновике
	StatusDraft ApplicationStatus = "draft"
	// StatusSubmitted заявление отправлено
	StatusSubmitted ApplicationStatus = "submitted"
	// StatusReview заявление на рассмотрении
	StatusReview ApplicationStatus = "review"
	// StatusApproved заявление одобрено
	StatusApproved ApplicationStatus = "approved"
	// StatusRejected заявление отклонено
	StatusRejected ApplicationStatus = "rejected"
)

// Application представляет заявление абитуриента
type Application struct {
	ID                 int              `json:"id"`
	ApplicantID        int              `json:"applicant_id"`
	Faculty            string           `json:"faculty"`
	Program            string           `json:"program"`
	Status             ApplicationStatus `json:"status"`
	DocumentsSubmitted bool             `json:"documents_submitted"`
	PersonalInfo       PersonalInfo     `json:"personal_info" gorm:"embedded"`
	EducationInfo      EducationInfo    `json:"education_info" gorm:"embedded"`
	Comments           string           `json:"comments"`
	CreatedAt          time.Time        `json:"created_at"`
	UpdatedAt          time.Time        `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Application
func (a *Application) TableName() string {
	return "applicant.applications"
}

// PersonalInfo содержит персональные данные абитуриента
type PersonalInfo struct {
	PassportSeries   string    `json:"passport_series"`
	PassportNumber   string    `json:"passport_number"`
	PassportIssuedBy string    `json:"passport_issued_by"`
	PassportDate     time.Time `json:"passport_date"`
	BirthDate        time.Time `json:"birth_date"`
	BirthPlace       string    `json:"birth_place"`
	Address          string    `json:"address"`
}

// EducationInfo содержит информацию об образовании абитуриента
type EducationInfo struct {
	EducationType        string  `json:"education_type"`    // Среднее, среднее специальное, высшее
	Institution          string  `json:"institution"`       // Название учебного заведения
	GraduationYear       int     `json:"graduation_year"`
	DocumentNumber       string  `json:"document_number"`   // Номер аттестата/диплома
	DocumentDate         string  `json:"document_date"`     // Дата выдачи документа
	AverageGrade         float32 `json:"average_grade"`
	HasOriginalDocuments bool    `json:"has_original_documents"`
}

// ApplicationRequest представляет запрос на создание или обновление заявления
type ApplicationRequest struct {
	Faculty            string       `json:"faculty" binding:"required"`
	Program            string       `json:"program" binding:"required"`
	PersonalInfo       PersonalInfo `json:"personal_info" binding:"required"`
	EducationInfo      EducationInfo `json:"education_info" binding:"required"`
	Comments           string       `json:"comments"`
}



