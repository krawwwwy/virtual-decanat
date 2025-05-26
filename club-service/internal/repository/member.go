package repository

import (
	"club-service/internal/model"
	"database/sql"
)

type MemberRepositoryInterface interface {
	Add(member *model.Member) error
	GetByClubID(clubID int) ([]model.Member, error)
	Remove(id int) error
	GetByUserID(userID int) ([]model.Member, error)
}

type MemberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

func (r *MemberRepository) Add(member *model.Member) error {
	query := `INSERT INTO club_members (club_id, user_id, role) VALUES ($1, $2, $3) RETURNING id, joined_at`
	return r.db.QueryRow(query, member.ClubID, member.UserID, member.Role).Scan(&member.ID, &member.JoinedAt)
}

func (r *MemberRepository) GetByClubID(clubID int) ([]model.Member, error) {
	query := `SELECT id, club_id, user_id, role, joined_at FROM club_members WHERE club_id = $1`
	rows, err := r.db.Query(query, clubID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []model.Member
	for rows.Next() {
		var member model.Member
		if err := rows.Scan(&member.ID, &member.ClubID, &member.UserID, &member.Role, &member.JoinedAt); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (r *MemberRepository) Remove(id int) error {
	query := `DELETE FROM club_members WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *MemberRepository) GetByUserID(userID int) ([]model.Member, error) {
	query := `SELECT id, club_id, user_id, role, joined_at FROM club_members WHERE user_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []model.Member
	for rows.Next() {
		var member model.Member
		if err := rows.Scan(&member.ID, &member.ClubID, &member.UserID, &member.Role, &member.JoinedAt); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
} 