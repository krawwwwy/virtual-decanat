package service

import (
	"context"
	"performance-service/internal/model"
	"performance-service/internal/repository"
)

type PerformanceService struct {
	repo repository.PerformanceRepository
}

func NewPerformanceService(repo repository.PerformanceRepository) *PerformanceService {
	return &PerformanceService{repo: repo}
}

func (s *PerformanceService) GetPerformance(ctx context.Context, studentID string) (*model.StudentPerformance, error) {
	return s.repo.GetPerformance(ctx, studentID)
}

func (s *PerformanceService) GetGrades(ctx context.Context, studentID string) ([]model.Grade, error) {
	return s.repo.GetGrades(ctx, studentID)
}

func (s *PerformanceService) GetAttendance(ctx context.Context, studentID string) ([]model.Attendance, error) {
	return s.repo.GetAttendance(ctx, studentID)
}

func (s *PerformanceService) GetDebts(ctx context.Context, studentID string) ([]model.Debt, error) {
	return s.repo.GetDebts(ctx, studentID)
}

func (s *PerformanceService) GetRating(ctx context.Context, studentID string) (float64, error) {
	return s.repo.GetRating(ctx, studentID)
} 