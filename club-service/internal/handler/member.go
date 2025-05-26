package handler

import (
	"club-service/internal/model"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MemberServiceInterface interface {
	AddMember(member *model.Member) error
	GetClubMembers(clubID int) ([]model.Member, error)
	RemoveMember(id int) error
	GetUserClubs(userID int) ([]model.Member, error)
}

type MemberHandler struct {
	service MemberServiceInterface
}

func NewMemberHandler(service MemberServiceInterface) *MemberHandler {
	return &MemberHandler{service: service}
}

func (h *MemberHandler) AddMember(c *gin.Context) {
	var member model.Member
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

func (h *MemberHandler) DebugAddMember(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var member model.Member
	err = json.Unmarshal(body, &member)
	c.JSON(http.StatusOK, gin.H{
		"raw_body": string(body),
		"unmarshal_error": err,
		"parsed_member": member,
	})
}

func (h *MemberHandler) GetUserClubs(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	clubs, err := h.service.GetUserClubs(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clubs)
} 