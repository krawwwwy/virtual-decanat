package service

import (
	"club-service/internal/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMemberRepo struct {
	mock.Mock
}

func (m *MockMemberRepo) Add(member *repository.Member) error {
	args := m.Called(member)
	return args.Error(0)
}
func (m *MockMemberRepo) GetByClubID(clubID int) ([]repository.Member, error) {
	args := m.Called(clubID)
	return args.Get(0).([]repository.Member), args.Error(1)
}
func (m *MockMemberRepo) Remove(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockMemberRepo) GetByUserID(userID int) ([]repository.Member, error) {
	args := m.Called(userID)
	return args.Get(0).([]repository.Member), args.Error(1)
}

func TestMemberService_AddMember(t *testing.T) {
	repo := new(MockMemberRepo)
	service := NewMemberService(repo)
	member := &repository.Member{ClubID: 1, UserID: 2, Role: "admin"}
	repo.On("Add", member).Return(nil)

	err := service.AddMember(member)
	assert.NoError(t, err)
}

func TestMemberService_GetClubMembers(t *testing.T) {
	repo := new(MockMemberRepo)
	service := NewMemberService(repo)
	members := []repository.Member{{ID: 1, ClubID: 1, UserID: 2, Role: "admin"}}
	repo.On("GetByClubID", 1).Return(members, nil)

	result, err := service.GetClubMembers(1)
	assert.NoError(t, err)
	assert.Equal(t, members, result)
}

func TestMemberService_RemoveMember_NotFound(t *testing.T) {
	repo := new(MockMemberRepo)
	service := NewMemberService(repo)
	repo.On("Remove", 42).Return(errors.New("not found"))

	err := service.RemoveMember(42)
	assert.Error(t, err)
} 