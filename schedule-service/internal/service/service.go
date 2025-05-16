package service

import (
	"context"

	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/model"
	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/repository"
)

type ScheduleService interface {
	CreateSchedule(ctx context.Context, req model.CreateScheduleRequest) error
	UpdateSchedule(ctx context.Context, id uint, req model.UpdateScheduleRequest) error
	DeleteSchedule(ctx context.Context, id uint) error
	GetScheduleByID(ctx context.Context, id uint) (*model.ScheduleResponse, error)
	ListByTeacher(ctx context.Context, teacherID uint) ([]model.ScheduleResponse, error)
	ListByGroup(ctx context.Context, groupID uint) ([]model.ScheduleResponse, error)

	CreateGroup(ctx context.Context, req model.CreateGroupRequest) error
	UpdateGroup(ctx context.Context, id uint, req model.UpdateGroupRequest) error
	DeleteGroup(ctx context.Context, id uint) error
	GetGroupByID(ctx context.Context, id uint) (*model.Group, error)
	ListGroups(ctx context.Context) ([]model.Group, error)

	CreateSubject(ctx context.Context, req model.CreateSubjectRequest) error
	UpdateSubject(ctx context.Context, id uint, req model.UpdateSubjectRequest) error
	DeleteSubject(ctx context.Context, id uint) error
	GetSubjectByID(ctx context.Context, id uint) (*model.Subject, error)
	ListSubjects(ctx context.Context) ([]model.Subject, error)
}

type ScheduleServiceImpl struct {
	repo repository.ScheduleRepository
}

func NewScheduleService(repo repository.ScheduleRepository) ScheduleService {
	return &ScheduleServiceImpl{repo: repo}
}

func (s *ScheduleServiceImpl) CreateSchedule(ctx context.Context, req model.CreateScheduleRequest) error {
	sch := &model.Schedule{
		SubjectID:  req.SubjectID,
		TeacherID: req.TeacherID,
		GroupID:   req.GroupID,
		DayOfWeek: req.DayOfWeek,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Room:      req.Room,
	}
	return s.repo.CreateSchedule(ctx, sch)
}
func (s *ScheduleServiceImpl) UpdateSchedule(ctx context.Context, id uint, req model.UpdateScheduleRequest) error {
	sch, err := s.repo.GetScheduleByID(ctx, id)
	if err != nil {
		return err
	}
	if req.SubjectID != 0 {
		sch.SubjectID = req.SubjectID
	}
	if req.TeacherID != 0 {
		sch.TeacherID = req.TeacherID
	}
	if req.GroupID != 0 {
		sch.GroupID = req.GroupID
	}
	if req.DayOfWeek != 0 {
		sch.DayOfWeek = req.DayOfWeek
	}
	if !req.StartTime.IsZero() {
		sch.StartTime = req.StartTime
	}
	if !req.EndTime.IsZero() {
		sch.EndTime = req.EndTime
	}
	if req.Room != "" {
		sch.Room = req.Room
	}
	return s.repo.UpdateSchedule(ctx, sch)
}
func (s *ScheduleServiceImpl) DeleteSchedule(ctx context.Context, id uint) error {
	return s.repo.DeleteSchedule(ctx, id)
}
func (s *ScheduleServiceImpl) GetScheduleByID(ctx context.Context, id uint) (*model.ScheduleResponse, error) {
	sch, err := s.repo.GetScheduleByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toScheduleResponse(sch), nil
}
func (s *ScheduleServiceImpl) ListByTeacher(ctx context.Context, teacherID uint) ([]model.ScheduleResponse, error) {
	schs, err := s.repo.ListByTeacher(ctx, teacherID)
	if err != nil {
		return nil, err
	}
	return toScheduleResponses(schs), nil
}
func (s *ScheduleServiceImpl) ListByGroup(ctx context.Context, groupID uint) ([]model.ScheduleResponse, error) {
	schs, err := s.repo.ListByGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return toScheduleResponses(schs), nil
}
// --- Group ---
func (s *ScheduleServiceImpl) CreateGroup(ctx context.Context, req model.CreateGroupRequest) error {
	g := &model.Group{
		Name:    req.Name,
		Faculty: req.Faculty,
		Year:    req.Year,
	}
	return s.repo.CreateGroup(ctx, g)
}
func (s *ScheduleServiceImpl) UpdateGroup(ctx context.Context, id uint, req model.UpdateGroupRequest) error {
	g, err := s.repo.GetGroupByID(ctx, id)
	if err != nil {
		return err
	}
	if req.Name != "" {
		g.Name = req.Name
	}
	if req.Faculty != "" {
		g.Faculty = req.Faculty
	}
	if req.Year != 0 {
		g.Year = req.Year
	}
	return s.repo.UpdateGroup(ctx, g)
}
func (s *ScheduleServiceImpl) DeleteGroup(ctx context.Context, id uint) error {
	return s.repo.DeleteGroup(ctx, id)
}
func (s *ScheduleServiceImpl) GetGroupByID(ctx context.Context, id uint) (*model.Group, error) {
	return s.repo.GetGroupByID(ctx, id)
}
func (s *ScheduleServiceImpl) ListGroups(ctx context.Context) ([]model.Group, error) {
	return s.repo.ListGroups(ctx)
}
// --- Subject ---
func (s *ScheduleServiceImpl) CreateSubject(ctx context.Context, req model.CreateSubjectRequest) error {
	sub := &model.Subject{
		Name:    req.Name,
		Code:    req.Code,
		Credits: req.Credits,
	}
	return s.repo.CreateSubject(ctx, sub)
}
func (s *ScheduleServiceImpl) UpdateSubject(ctx context.Context, id uint, req model.UpdateSubjectRequest) error {
	sub, err := s.repo.GetSubjectByID(ctx, id)
	if err != nil {
		return err
	}
	if req.Name != "" {
		sub.Name = req.Name
	}
	if req.Code != "" {
		sub.Code = req.Code
	}
	if req.Credits != 0 {
		sub.Credits = req.Credits
	}
	return s.repo.UpdateSubject(ctx, sub)
}
func (s *ScheduleServiceImpl) DeleteSubject(ctx context.Context, id uint) error {
	return s.repo.DeleteSubject(ctx, id)
}
func (s *ScheduleServiceImpl) GetSubjectByID(ctx context.Context, id uint) (*model.Subject, error) {
	return s.repo.GetSubjectByID(ctx, id)
}
func (s *ScheduleServiceImpl) ListSubjects(ctx context.Context) ([]model.Subject, error) {
	return s.repo.ListSubjects(ctx)
}
// --- helpers ---
func toScheduleResponse(s *model.Schedule) *model.ScheduleResponse {
	return &model.ScheduleResponse{
		ID:        s.ID,
		Subject:   s.Subject.Name,
		SubjectID: s.SubjectID,
		Teacher:   "", // TODO: teacher name if needed
		TeacherID: s.TeacherID,
		Group:     s.Group.Name,
		GroupID:   s.GroupID,
		DayOfWeek: s.DayOfWeek,
		StartTime: s.StartTime,
		EndTime:   s.EndTime,
		Room:      s.Room,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}
func toScheduleResponses(list []model.Schedule) []model.ScheduleResponse {
	res := make([]model.ScheduleResponse, 0, len(list))
	for _, s := range list {
		res = append(res, *toScheduleResponse(&s))
	}
	return res
} 