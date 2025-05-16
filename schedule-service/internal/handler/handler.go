package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/model"
	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/service"
)

type ScheduleHandler struct {
	service service.ScheduleService
}

func NewScheduleHandler(s service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{service: s}
}

// --- Schedule ---
func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var req model.CreateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := h.service.CreateSchedule(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.SuccessResponse{Message: "Schedule created"})
}
func (h *ScheduleHandler) UpdateSchedule(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req model.UpdateScheduleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := h.service.UpdateSchedule(c.Request.Context(), uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Schedule updated"})
}
func (h *ScheduleHandler) DeleteSchedule(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.DeleteSchedule(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Schedule deleted"})
}
func (h *ScheduleHandler) GetScheduleByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := h.service.GetScheduleByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *ScheduleHandler) ListByTeacher(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("teacher_id"), 10, 64)
	res, err := h.service.ListByTeacher(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *ScheduleHandler) ListByGroup(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("group_id"), 10, 64)
	res, err := h.service.ListByGroup(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
// --- Group ---
func (h *ScheduleHandler) CreateGroup(c *gin.Context) {
	var req model.CreateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := h.service.CreateGroup(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.SuccessResponse{Message: "Group created"})
}
func (h *ScheduleHandler) UpdateGroup(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req model.UpdateGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := h.service.UpdateGroup(c.Request.Context(), uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Group updated"})
}
func (h *ScheduleHandler) DeleteGroup(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.DeleteGroup(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Group deleted"})
}
func (h *ScheduleHandler) GetGroupByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := h.service.GetGroupByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *ScheduleHandler) ListGroups(c *gin.Context) {
	res, err := h.service.ListGroups(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
// --- Subject ---
func (h *ScheduleHandler) CreateSubject(c *gin.Context) {
	var req model.CreateSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := h.service.CreateSubject(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.SuccessResponse{Message: "Subject created"})
}
func (h *ScheduleHandler) UpdateSubject(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var req model.UpdateSubjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}
	if err := h.service.UpdateSubject(c.Request.Context(), uint(id), req); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Subject updated"})
}
func (h *ScheduleHandler) DeleteSubject(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if err := h.service.DeleteSubject(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Subject deleted"})
}
func (h *ScheduleHandler) GetSubjectByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	res, err := h.service.GetSubjectByID(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
func (h *ScheduleHandler) ListSubjects(c *gin.Context) {
	res, err := h.service.ListSubjects(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
} 