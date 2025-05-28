package model

import "time"

// Group представляет группу студентов
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