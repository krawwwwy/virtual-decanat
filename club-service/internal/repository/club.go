package repository

import (
	"database/sql"
	"errors"
)

var ErrNotFound = errors.New("not found")

type Club struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type ClubRepositoryInterface interface {
	Create(club *Club) error
	GetAll() ([]Club, error)
	GetByID(id int) (*Club, error)
	Update(club *Club) error
	Delete(id int) error
}

type ClubRepository struct {
	db *sql.DB
}

func NewClubRepository(db *sql.DB) *ClubRepository {
	return &ClubRepository{db: db}
}

func (r *ClubRepository) Create(club *Club) error {
	query := `INSERT INTO club.clubs (name, description) VALUES ($1, $2) RETURNING id, created_at`
	return r.db.QueryRow(query, club.Name, club.Description).Scan(&club.ID, &club.CreatedAt)
}

func (r *ClubRepository) GetAll() ([]Club, error) {
	query := `SELECT id, name, description, created_at FROM club.clubs`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clubs []Club
	for rows.Next() {
		var club Club
		if err := rows.Scan(&club.ID, &club.Name, &club.Description, &club.CreatedAt); err != nil {
			return nil, err
		}
		clubs = append(clubs, club)
	}
	return clubs, nil
}

func (r *ClubRepository) GetByID(id int) (*Club, error) {
	query := `SELECT id, name, description, created_at FROM club.clubs WHERE id = $1`
	club := &Club{}
	err := r.db.QueryRow(query, id).Scan(&club.ID, &club.Name, &club.Description, &club.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return club, nil
}

func (r *ClubRepository) Update(club *Club) error {
	query := `UPDATE club.clubs SET name = $1, description = $2 WHERE id = $3`
	result, err := r.db.Exec(query, club.Name, club.Description, club.ID)
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

func (r *ClubRepository) Delete(id int) error {
	query := `DELETE FROM club.clubs WHERE id = $1`
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