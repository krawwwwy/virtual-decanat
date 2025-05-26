package service

import (
	"club-service/internal/model"
	"club-service/internal/repository"
)

type MemberService struct {
	repo repository.MemberRepositoryInterface
}

func NewMemberService(repo repository.MemberRepositoryInterface) *MemberService {
	return &MemberService{repo: repo}
}

func (s *MemberService) AddMember(member *model.Member) error {
	return s.repo.Add(member)
}

func (s *MemberService) GetClubMembers(clubID int) ([]model.Member, error) {
	return s.repo.GetByClubID(clubID)
}

func (s *MemberService) RemoveMember(id int) error {
	return s.repo.Remove(id)
}

func (s *MemberService) GetUserClubs(userID int) ([]model.Member, error) {
	return s.repo.GetByUserID(userID)
} 