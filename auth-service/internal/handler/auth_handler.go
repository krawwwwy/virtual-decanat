package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/model"
	"github.com/krawwwwy/virtual-decanat/auth-service/internal/service"
	"go.uber.org/zap"
)

// AuthHandler обрабатывает запросы аутентификации
type AuthHandler struct {
	authService service.AuthService
	logger      *zap.Logger
}

// NewAuthHandler создает новый AuthHandler
func NewAuthHandler(authService service.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger:      logger,
	}
}

// Register обрабатывает запрос на регистрацию
func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}

	user, err := h.authService.Register(c.Request.Context(), req)
	if err != nil {
		h.logger.Error("Failed to register user", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login обрабатывает запрос на вход
func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}

	tokens, err := h.authService.Login(c.Request.Context(), req)
	if err != nil {
		h.logger.Error("Failed to login", zap.Error(err))
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// RefreshToken обрабатывает запрос на обновление токена
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req model.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}

	tokens, err := h.authService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		h.logger.Error("Failed to refresh token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

// GetProfile обрабатывает запрос на получение профиля
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		h.logger.Error("Failed to get user ID from context", zap.Error(err))
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Unauthorized"})
		return
	}

	user, err := h.authService.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		h.logger.Error("Failed to get user profile", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateProfile обрабатывает запрос на обновление профиля
func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		h.logger.Error("Failed to get user ID from context", zap.Error(err))
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Unauthorized"})
		return
	}

	var req model.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}

	user, err := h.authService.UpdateProfile(c.Request.Context(), userID, req)
	if err != nil {
		h.logger.Error("Failed to update profile", zap.Error(err))
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ChangePassword обрабатывает запрос на изменение пароля
func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userID, err := getUserIDFromContext(c)
	if err != nil {
		h.logger.Error("Failed to get user ID from context", zap.Error(err))
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "Unauthorized"})
		return
	}

	var req model.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if err := h.authService.ChangePassword(c.Request.Context(), userID, req); err != nil {
		h.logger.Error("Failed to change password", zap.Error(err))
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.SuccessResponse{Message: "Password changed successfully"})
}

// getUserIDFromContext получает ID пользователя из контекста
func getUserIDFromContext(c *gin.Context) (uint, error) {
	userIDStr, exists := c.Get("userID")
	if !exists {
		return 0, http.ErrNoCookie
	}

	switch id := userIDStr.(type) {
	case string:
		userID, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(userID), nil
	case float64:
		return uint(id), nil
	case uint:
		return id, nil
	default:
		return 0, http.ErrNoCookie
	}
} 