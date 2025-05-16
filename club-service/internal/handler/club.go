package handler

import (
	"club-service/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClubServiceInterface interface {
	CreateClub(club *repository.Club) error
	GetAllClubs() ([]repository.Club, error)
	GetClubByID(id int) (*repository.Club, error)
	UpdateClub(club *repository.Club) error
	DeleteClub(id int) error
}

type ClubHandler struct {
	service ClubServiceInterface
}

func NewClubHandler(service ClubServiceInterface) *ClubHandler {
	return &ClubHandler{service: service}
}

func (h *ClubHandler) CreateClub(c *gin.Context) {
	var club repository.Club
	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateClub(&club); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, club)
}

func (h *ClubHandler) GetAllClubs(c *gin.Context) {
	clubs, err := h.service.GetAllClubs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clubs)
}

func (h *ClubHandler) GetClubByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	club, err := h.service.GetClubByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, club)
}

func (h *ClubHandler) UpdateClub(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var club repository.Club
	if err := c.ShouldBindJSON(&club); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	club.ID = id
	if err := h.service.UpdateClub(&club); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, club)
}

func (h *ClubHandler) DeleteClub(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.DeleteClub(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
} 