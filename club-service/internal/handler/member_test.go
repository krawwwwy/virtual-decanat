package handler

import (
	"bytes"
	"club-service/internal/repository"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockMemberService struct {
	members []repository.Member
}

func (m *mockMemberService) AddMember(member *repository.Member) error {
	member.ID = len(m.members) + 1
	m.members = append(m.members, *member)
	return nil
}
func (m *mockMemberService) GetClubMembers(clubID int) ([]repository.Member, error) {
	var res []repository.Member
	for _, mem := range m.members {
		if mem.ClubID == clubID {
			res = append(res, mem)
		}
	}
	return res, nil
}
func (m *mockMemberService) RemoveMember(id int) error {
	for i, mem := range m.members {
		if mem.ID == id {
			m.members = append(m.members[:i], m.members[i+1:]...)
			return nil
		}
	}
	return nil
}
func (m *mockMemberService) GetUserClubs(userID int) ([]repository.Member, error) {
	return nil, nil
}

func setupRouterMember() (*gin.Engine, *mockMemberService) {
	gin.SetMode(gin.TestMode)
	mock := &mockMemberService{}
	h := NewMemberHandler(mock)
	r := gin.Default()
	r.POST("/api/members/", h.AddMember)
	r.GET("/api/members/club/:clubId", h.GetClubMembers)
	r.DELETE("/api/members/:id", h.RemoveMember)
	return r, mock
}

func TestAddMemberHandler(t *testing.T) {
	r, _ := setupRouterMember()
	body := map[string]interface{}{"club_id": 1, "user_id": 2, "role": "admin"}
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/members/", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetClubMembersHandler(t *testing.T) {
	r, mock := setupRouterMember()
	mock.members = []repository.Member{{ID: 1, ClubID: 1, UserID: 2, Role: "admin"}}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/members/club/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRemoveMemberHandler(t *testing.T) {
	r, mock := setupRouterMember()
	mock.members = []repository.Member{{ID: 1, ClubID: 1, UserID: 2, Role: "admin"}}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/members/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestAddMemberHandler_Integration(t *testing.T) {
	r, _ := setupRouterMember()
	body := map[string]interface{}{ "club_id": 1, "user_id": 1, "role": "integration_test" }
	jsonBody, _ := json.Marshal(body)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/members/", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
} 