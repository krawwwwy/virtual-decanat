package service

import (
	"club-service/internal/model"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClubRepo struct {
	mock.Mock
}

func (m *MockClubRepo) Create(club *model.Club) error {
	args := m.Called(club)
	return args.Error(0)
}
func (m *MockClubRepo) GetAll() ([]model.Club, error) {
	args := m.Called()
	return args.Get(0).([]model.Club), args.Error(1)
}
func (m *MockClubRepo) GetByID(id int) (*model.Club, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Club), args.Error(1)
}
func (m *MockClubRepo) Update(club *model.Club) error {
	args := m.Called(club)
	return args.Error(0)
}
func (m *MockClubRepo) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestClubService_CreateClub(t *testing.T) {
	repo := new(MockClubRepo)
	service := NewClubService(repo)
	club := &model.Club{Name: "Test", Description: "Desc"}
	repo.On("Create", club).Return(nil)

	err := service.CreateClub(club)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestClubService_GetAllClubs(t *testing.T) {
	repo := new(MockClubRepo)
	service := NewClubService(repo)
	clubs := []model.Club{{ID: 1, Name: "A"}}
	repo.On("GetAll").Return(clubs, nil)

	result, err := service.GetAllClubs()
	assert.NoError(t, err)
	assert.Equal(t, clubs, result)
}

func TestClubService_GetClubByID_NotFound(t *testing.T) {
	repo := new(MockClubRepo)
	service := NewClubService(repo)
	repo.On("GetByID", 42).Return(&model.Club{}, errors.New("not found"))

	_, err := service.GetClubByID(42)
	assert.Error(t, err)
}

func TestClubService_UpdateClub(t *testing.T) {
	repo := new(MockClubRepo)
	service := NewClubService(repo)
	club := &model.Club{ID: 1, Name: "Upd"}
	repo.On("Update", club).Return(nil)

	err := service.UpdateClub(club)
	assert.NoError(t, err)
}

func TestClubService_DeleteClub(t *testing.T) {
	repo := new(MockClubRepo)
	service := NewClubService(repo)
	repo.On("Delete", 1).Return(nil)

	err := service.DeleteClub(1)
	assert.NoError(t, err)
} 