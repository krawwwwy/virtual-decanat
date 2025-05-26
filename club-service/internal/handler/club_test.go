package handler

import (
	"bytes"
	"club-service/internal/model"
	"club-service/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockClubService struct {
	clubs []model.Club
}

func (m *mockClubService) CreateClub(club *model.Club) error {
	club.ID = len(m.clubs) + 1
	m.clubs = append(m.clubs, *club)
	return nil
}
func (m *mockClubService) GetAllClubs() ([]model.Club, error) {
	return m.clubs, nil
}
func (m *mockClubService) GetClubByID(id int) (*model.Club, error) {
	for _, c := range m.clubs {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, service.ErrNotFound
}
func (m *mockClubService) UpdateClub(club *model.Club) error {
	for i, c := range m.clubs {
		if c.ID == club.ID {
			m.clubs[i] = *club
			return nil
		}
	}
	return service.ErrNotFound
}
func (m *mockClubService) DeleteClub(id int) error {
	for i, c := range m.clubs {
		if c.ID == id {
			m.clubs = append(m.clubs[:i], m.clubs[i+1:]...)
			return nil
		}
	}
	return service.ErrNotFound
}

func setupRouterClub() (*gin.Engine, *mockClubService) {
	gin.SetMode(gin.TestMode)
	mock := &mockClubService{}
	h := NewClubHandler(mock)
	r := gin.Default()
	r.POST("/api/clubs/", h.CreateClub)
	r.GET("/api/clubs/", h.GetAllClubs)
	r.GET("/api/clubs/:id", h.GetClubByID)
	r.PUT("/api/clubs/:id", h.UpdateClub)
	r.DELETE("/api/clubs/:id", h.DeleteClub)
	return r, mock
}

func TestCreateClubHandler(t *testing.T) {
	r, _ := setupRouterClub()
	body := map[string]interface{}{"name": "Chess", "description": "Chess club"}
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/clubs/", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetAllClubsHandler(t *testing.T) {
	r, mock := setupRouterClub()
	mock.clubs = []model.Club{{ID: 1, Name: "A"}}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/clubs/", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetClubByID_NotFound(t *testing.T) {
	r, _ := setupRouterClub()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/clubs/42", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
} 