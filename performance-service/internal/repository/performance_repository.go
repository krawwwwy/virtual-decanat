package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"performance-service/internal/model"
)

type PerformanceRepository interface {
	GetPerformance(ctx context.Context, studentID string) (*model.StudentPerformance, error)
	GetGrades(ctx context.Context, studentID string) ([]model.Grade, error)
	GetAttendance(ctx context.Context, studentID string) ([]model.Attendance, error)
	GetDebts(ctx context.Context, studentID string) ([]model.Debt, error)
	GetRating(ctx context.Context, studentID string) (float64, error)
}

type PgxPerformanceRepository struct {
	db *pgxpool.Pool
}

func NewPgxPerformanceRepository(db *pgxpool.Pool) *PgxPerformanceRepository {
	return &PgxPerformanceRepository{db: db}
}

func (r *PgxPerformanceRepository) GetGrades(ctx context.Context, studentID string) ([]model.Grade, error) {
	rows, err := r.db.Query(ctx, `SELECT s.name as discipline, g.grade FROM performance.grades g JOIN schedule.subjects s ON g.subject_id = s.id WHERE g.student_id = $1`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var grades []model.Grade
	for rows.Next() {
		var g model.Grade
		if err := rows.Scan(&g.Discipline, &g.Value); err != nil {
			return nil, err
		}
		grades = append(grades, g)
	}
	return grades, nil
}

func (r *PgxPerformanceRepository) GetAttendance(ctx context.Context, studentID string) ([]model.Attendance, error) {
	rows, err := r.db.Query(ctx, `SELECT s.name as discipline, COUNT(a.id) as total, SUM(CASE WHEN a.is_present THEN 1 ELSE 0 END) as attended FROM performance.attendance a JOIN schedule.schedule sch ON a.schedule_id = sch.id JOIN schedule.subjects s ON sch.subject_id = s.id WHERE a.student_id = $1 GROUP BY s.name`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var att []model.Attendance
	for rows.Next() {
		var a model.Attendance
		if err := rows.Scan(&a.Discipline, &a.Total, &a.Attended); err != nil {
			return nil, err
		}
		att = append(att, a)
	}
	return att, nil
}

func (r *PgxPerformanceRepository) GetDebts(ctx context.Context, studentID string) ([]model.Debt, error) {
	rows, err := r.db.Query(ctx, `SELECT s.name as discipline, d.description FROM performance.debts d JOIN schedule.subjects s ON d.subject_id = s.id WHERE d.student_id = $1`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var debts []model.Debt
	for rows.Next() {
		var d model.Debt
		if err := rows.Scan(&d.Discipline, &d.Reason); err != nil {
			return nil, err
		}
		debts = append(debts, d)
	}
	return debts, nil
}

func (r *PgxPerformanceRepository) GetRating(ctx context.Context, studentID string) (float64, error) {
	var rating float64
	err := r.db.QueryRow(ctx, `SELECT COALESCE(AVG(grade),0) FROM performance.grades WHERE student_id = $1`, studentID).Scan(&rating)
	return rating, err
}

func (r *PgxPerformanceRepository) GetPerformance(ctx context.Context, studentID string) (*model.StudentPerformance, error) {
	grades, err := r.GetGrades(ctx, studentID)
	if err != nil {
		return nil, err
	}
	att, err := r.GetAttendance(ctx, studentID)
	if err != nil {
		return nil, err
	}
	debts, err := r.GetDebts(ctx, studentID)
	if err != nil {
		return nil, err
	}
	rating, err := r.GetRating(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return &model.StudentPerformance{
		StudentID:  studentID,
		Grades:     grades,
		Attendance: att,
		Debts:      debts,
		Rating:     rating,
	}, nil
} 