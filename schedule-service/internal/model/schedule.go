package model

import (
	"time"
)

// Group представляет учебную группу
type Group struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null"`
	Faculty   string    `json:"faculty" gorm:"not null"`
	Year      int       `json:"year" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Group
func (Group) TableName() string {
	return "schedule.groups"
}

// Subject представляет учебный предмет
type Subject struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	Code      string    `json:"code" gorm:"unique;not null"`
	Credits   int       `json:"credits" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Subject
func (Subject) TableName() string {
	return "schedule.subjects"
}

// Schedule представляет запись в расписании
type Schedule struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	SubjectID  uint      `json:"subject_id" gorm:"not null"`
	Subject    *Subject  `json:"subject,omitempty" gorm:"foreignKey:SubjectID"`
	TeacherID  uint      `json:"teacher_id" gorm:"not null"`
	GroupID    uint      `json:"group_id" gorm:"not null"`
	Group      *Group    `json:"group,omitempty" gorm:"foreignKey:GroupID"`
	DayOfWeek  int       `json:"day_of_week" gorm:"not null"` // 1 - Понедельник, 2 - Вторник и т.д.
	StartTime  time.Time `json:"start_time" gorm:"not null"`
	EndTime    time.Time `json:"end_time" gorm:"not null"`
	Room       string    `json:"room" gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// TableName указывает имя таблицы для модели Schedule
func (Schedule) TableName() string {
	return "schedule.schedule"
}

// CreateScheduleRequest представляет запрос на создание записи в расписании
type CreateScheduleRequest struct {
	SubjectID uint      `json:"subject_id" binding:"required"`
	TeacherID uint      `json:"teacher_id" binding:"required"`
	GroupID   uint      `json:"group_id" binding:"required"`
	DayOfWeek int       `json:"day_of_week" binding:"required,min=1,max=7"`
	StartTime time.Time `json:"start_time" binding:"required"`
	EndTime   time.Time `json:"end_time" binding:"required"`
	Room      string    `json:"room" binding:"required"`
}

// UpdateScheduleRequest представляет запрос на обновление записи в расписании
type UpdateScheduleRequest struct {
	SubjectID uint      `json:"subject_id"`
	TeacherID uint      `json:"teacher_id"`
	GroupID   uint      `json:"group_id"`
	DayOfWeek int       `json:"day_of_week" binding:"omitempty,min=1,max=7"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Room      string    `json:"room"`
}

// ScheduleResponse представляет ответ с данными расписания
type ScheduleResponse struct {
	ID         uint      `json:"id"`
	Subject    string    `json:"subject"`
	SubjectID  uint      `json:"subject_id"`
	Teacher    string    `json:"teacher"`
	TeacherID  uint      `json:"teacher_id"`
	Group      string    `json:"group"`
	GroupID    uint      `json:"group_id"`
	DayOfWeek  int       `json:"day_of_week"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Room       string    `json:"room"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateGroupRequest представляет запрос на создание группы
type CreateGroupRequest struct {
	Name    string `json:"name" binding:"required"`
	Faculty string `json:"faculty" binding:"required"`
	Year    int    `json:"year" binding:"required,min=1,max=6"`
}

// UpdateGroupRequest представляет запрос на обновление группы
type UpdateGroupRequest struct {
	Name    string `json:"name"`
	Faculty string `json:"faculty"`
	Year    int    `json:"year" binding:"omitempty,min=1,max=6"`
}

// CreateSubjectRequest представляет запрос на создание предмета
type CreateSubjectRequest struct {
	Name    string `json:"name" binding:"required"`
	Code    string `json:"code" binding:"required"`
	Credits int    `json:"credits" binding:"required,min=1"`
}

// UpdateSubjectRequest представляет запрос на обновление предмета
type UpdateSubjectRequest struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Credits int    `json:"credits" binding:"omitempty,min=1"`
}

// ErrorResponse представляет ответ с ошибкой
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse представляет успешный ответ
type SuccessResponse struct {
	Message string `json:"message"`
} 