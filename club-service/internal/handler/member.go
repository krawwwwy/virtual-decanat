package handler

import (
	"club-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MemberServiceInterface interface {
	AddMember(member *repository.Member) error
	GetClubMembers(clubID int) ([]repository.Member, error)
	RemoveMember(id int) error
	GetUserClubs(userID int) ([]repository.Member, error)
}

type MemberHandler struct {
	service MemberServiceInterface
}

func NewMemberHandler(service MemberServiceInterface) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) AddMember(c *gin.Context) {
	var member repository.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AddMember(&member); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, member)
}

func (h *MemberHandler) GetClubMembers(c *gin.Context) {
	clubID, err := strconv.Atoi(c.Param("clubId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid club id"})
		return
	}

	members, err := h.service.GetClubMembers(clubID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

func (h *MemberHandler) RemoveMember(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.RemoveMember(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
} 