package repository

import (
	"database/sql"
)

type Member struct {
	ID        int    `json:"id"`
	ClubID    int    `json:"club_id"`
	UserID    int    `json:"user_id"`
	Role      string `json:"role"`
	JoinedAt  string `json:"joined_at"`
}

type MemberRepositoryInterface interface {
	Add(member *Member) error
	GetByClubID(clubID int) ([]Member, error)
	Remove(id int) error
	GetByUserID(userID int) ([]Member, error)
}

type MemberRepository struct {
	db *sql.DB
}

func NewMemberRepository(db *sql.DB) *MemberRepository {
	return &MemberRepository{db: db}
}

func (r *MemberRepository) Add(member *Member) error {
	query := `INSERT INTO club.memberships (club_id, student_id, role) VALUES ($1, $2, $3) RETURNING id, joined_at`
	return r.db.QueryRow(query, member.ClubID, member.UserID, member.Role).Scan(&member.ID, &member.JoinedAt)
}

func (r *MemberRepository) GetByClubID(clubID int) ([]Member, error) {
	query := `SELECT id, club_id, student_id as user_id, role, joined_at FROM club.memberships WHERE club_id = $1`
	rows, err := r.db.Query(query, clubID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		if err := rows.Scan(&member.ID, &member.ClubID, &member.UserID, &member.Role, &member.JoinedAt); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}

func (r *MemberRepository) Remove(id int) error {
	query := `DELETE FROM club.memberships WHERE id = $1`
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

func (r *MemberRepository) GetByUserID(userID int) ([]Member, error) {
	query := `SELECT id, club_id, student_id as user_id, role, joined_at FROM club.memberships WHERE student_id = $1`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member
		if err := rows.Scan(&member.ID, &member.ClubID, &member.UserID, &member.Role, &member.JoinedAt); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
} 