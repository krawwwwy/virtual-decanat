package service

import (
	"club-service/internal/repository"
)

var ErrNotFound = repository.ErrNotFound

type ClubService struct {
	repo repository.ClubRepositoryInterface
}

func NewClubService(repo repository.ClubRepositoryInterface) *ClubService {
	return &ClubService{repo: repo}
}

func (s *ClubService) CreateClub(club *repository.Club) error {
	return s.repo.Create(club)
}

func (s *ClubService) GetAllClubs() ([]repository.Club, error) {
	return s.repo.GetAll()
}

func (s *ClubService) GetClubByID(id int) (*repository.Club, error) {
	return s.repo.GetByID(id)
}

func (s *ClubService) UpdateClub(club *repository.Club) error {
	return s.repo.Update(club)
}

func (s *ClubService) DeleteClub(id int) error {
	return s.repo.Delete(id)
} 