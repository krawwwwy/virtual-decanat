package repository

import (
	"context"

	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/model"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(ctx context.Context, s *model.Schedule) error
	UpdateSchedule(ctx context.Context, s *model.Schedule) error
	DeleteSchedule(ctx context.Context, id uint) error
	GetScheduleByID(ctx context.Context, id uint) (*model.Schedule, error)
	ListByTeacher(ctx context.Context, teacherID uint) ([]model.Schedule, error)
	ListByGroup(ctx context.Context, groupID uint) ([]model.Schedule, error)

	CreateGroup(ctx context.Context, g *model.Group) error
	UpdateGroup(ctx context.Context, g *model.Group) error
	DeleteGroup(ctx context.Context, id uint) error
	GetGroupByID(ctx context.Context, id uint) (*model.Group, error)
	ListGroups(ctx context.Context) ([]model.Group, error)

	CreateSubject(ctx context.Context, s *model.Subject) error
	UpdateSubject(ctx context.Context, s *model.Subject) error
	DeleteSubject(ctx context.Context, id uint) error
	GetSubjectByID(ctx context.Context, id uint) (*model.Subject, error)
	ListSubjects(ctx context.Context) ([]model.Subject, error)
}

type ScheduleRepositoryImpl struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{db: db}
}

func (r *ScheduleRepositoryImpl) CreateSchedule(ctx context.Context, s *model.Schedule) error {
	return r.db.WithContext(ctx).Create(s).Error
}
func (r *ScheduleRepositoryImpl) UpdateSchedule(ctx context.Context, s *model.Schedule) error {
	return r.db.WithContext(ctx).Save(s).Error
}
func (r *ScheduleRepositoryImpl) DeleteSchedule(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Schedule{}, id).Error
}
func (r *ScheduleRepositoryImpl) GetScheduleByID(ctx context.Context, id uint) (*model.Schedule, error) {
	var s model.Schedule
	err := r.db.WithContext(ctx).Preload("Subject").Preload("Group").First(&s, id).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}
func (r *ScheduleRepositoryImpl) ListByTeacher(ctx context.Context, teacherID uint) ([]model.Schedule, error) {
	var res []model.Schedule
	err := r.db.WithContext(ctx).Preload("Subject").Preload("Group").Where("teacher_id = ?", teacherID).Find(&res).Error
	return res, err
}
func (r *ScheduleRepositoryImpl) ListByGroup(ctx context.Context, groupID uint) ([]model.Schedule, error) {
	var res []model.Schedule
	err := r.db.WithContext(ctx).Preload("Subject").Preload("Group").Where("group_id = ?", groupID).Find(&res).Error
	return res, err
}
// --- Group CRUD ---
func (r *ScheduleRepositoryImpl) CreateGroup(ctx context.Context, g *model.Group) error {
	return r.db.WithContext(ctx).Create(g).Error
}
func (r *ScheduleRepositoryImpl) UpdateGroup(ctx context.Context, g *model.Group) error {
	return r.db.WithContext(ctx).Save(g).Error
}
func (r *ScheduleRepositoryImpl) DeleteGroup(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Group{}, id).Error
}
func (r *ScheduleRepositoryImpl) GetGroupByID(ctx context.Context, id uint) (*model.Group, error) {
	var g model.Group
	err := r.db.WithContext(ctx).First(&g, id).Error
	if err != nil {
		return nil, err
	}
	return &g, nil
}
func (r *ScheduleRepositoryImpl) ListGroups(ctx context.Context) ([]model.Group, error) {
	var res []model.Group
	err := r.db.WithContext(ctx).Find(&res).Error
	return res, err
}
// --- Subject CRUD ---
func (r *ScheduleRepositoryImpl) CreateSubject(ctx context.Context, s *model.Subject) error {
	return r.db.WithContext(ctx).Create(s).Error
}
func (r *ScheduleRepositoryImpl) UpdateSubject(ctx context.Context, s *model.Subject) error {
	return r.db.WithContext(ctx).Save(s).Error
}
func (r *ScheduleRepositoryImpl) DeleteSubject(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Subject{}, id).Error
}
func (r *ScheduleRepositoryImpl) GetSubjectByID(ctx context.Context, id uint) (*model.Subject, error) {
	var s model.Subject
	err := r.db.WithContext(ctx).First(&s, id).Error
	if err != nil {
		return nil, err
	}
	return &s, nil
}
func (r *ScheduleRepositoryImpl) ListSubjects(ctx context.Context) ([]model.Subject, error) {
	var res []model.Subject
	err := r.db.WithContext(ctx).Find(&res).Error
	return res, err
} 