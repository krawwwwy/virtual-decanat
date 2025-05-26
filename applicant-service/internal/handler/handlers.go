package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/model"
	"github.com/krawwwwy/virtual-decanat/applicant-service/internal/service"
)

// Handler обработчик HTTP-запросов
type Handler struct {
	applicantService service.ApplicantService
}

// NewHandler создает новый Handler
func NewHandler(applicantService service.ApplicantService) *Handler {
	return &Handler{
		applicantService: applicantService,
	}
}

// SetupRoutes настраивает маршруты для Gin
func (h *Handler) SetupRoutes(r *gin.Engine) {
	// Группа путей API
	api := r.Group("/api/v1")
	
	// Публичные маршруты
	api.POST("/register", h.Register)
	api.POST("/login", h.Login)
	
	// Защищенные маршруты (требуется аутентификация)
	auth := api.Group("")
	auth.Use(h.AuthMiddleware())
	{
		// Маршруты для абитуриентов
		auth.GET("/applicant", h.GetApplicant)
		
		// Маршруты для заявлений
		applications := auth.Group("/applications")
		{
			applications.POST("", h.CreateApplication)
			applications.GET("", h.GetApplications)
			applications.GET("/:id", h.GetApplication)
			applications.PUT("/:id", h.UpdateApplication)
			applications.POST("/:id/submit", h.SubmitApplication)
			applications.GET("/:id/status", h.GetApplicationStatus)
		}
	}
}

// Register обрабатывает регистрацию нового абитуриента
func (h *Handler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	
	applicant, err := h.applicantService.Register(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, model.ErrEmailTaken) {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already taken"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register"})
		return
	}
	
	c.JSON(http.StatusCreated, applicant)
}

// Login обрабатывает вход в систему
func (h *Handler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	
	response, err := h.applicantService.Login(c.Request.Context(), &req)
	if err != nil {
		if errors.Is(err, model.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}
	
	c.JSON(http.StatusOK, response)
}

// GetApplicant возвращает информацию о текущем абитуриенте
func (h *Handler) GetApplicant(c *gin.Context) {
	applicantID := getApplicantID(c)
	
	applicant, err := h.applicantService.GetApplicant(c.Request.Context(), applicantID)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Applicant not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get applicant"})
		return
	}
	
	c.JSON(http.StatusOK, applicant)
}

// CreateApplication обрабатывает создание нового заявления
func (h *Handler) CreateApplication(c *gin.Context) {
	var req model.ApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	
	applicantID := getApplicantID(c)
	
	application, err := h.applicantService.CreateApplication(c.Request.Context(), applicantID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
		return
	}
	
	c.JSON(http.StatusCreated, application)
}

// GetApplications возвращает все заявления текущего абитуриента
func (h *Handler) GetApplications(c *gin.Context) {
	applicantID := getApplicantID(c)
	
	applications, err := h.applicantService.GetApplicationsByApplicantID(c.Request.Context(), applicantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch applications"})
		return
	}
	
	c.JSON(http.StatusOK, applications)
}

// GetApplication возвращает заявление по ID
func (h *Handler) GetApplication(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	
	application, err := h.applicantService.GetApplication(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get application"})
		return
	}
	
	// Проверяем, что заявление принадлежит текущему пользователю
	applicantID := getApplicantID(c)
	if uint(application.ApplicantID) != applicantID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}
	
	c.JSON(http.StatusOK, application)
}

// UpdateApplication обрабатывает обновление заявления
func (h *Handler) UpdateApplication(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	
	var req model.ApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
	
	application, err := h.applicantService.UpdateApplication(c.Request.Context(), uint(id), &req)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update application"})
		return
	}
	
	c.JSON(http.StatusOK, application)
}

// SubmitApplication обрабатывает отправку заявления
func (h *Handler) SubmitApplication(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	
	if err := h.applicantService.SubmitApplication(c.Request.Context(), uint(id)); err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit application"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"status": "submitted"})
}

// GetApplicationStatus возвращает статус заявления
func (h *Handler) GetApplicationStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	
	status, err := h.applicantService.GetApplicationStatus(c.Request.Context(), uint(id))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get application status"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// AuthMiddleware возвращает промежуточное ПО для аутентификации
func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}
		
		// Извлекаем токен из заголовка (формат: Bearer <token>)
		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		}
		
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			return
		}
		
		// В реальном приложении здесь будет проверка JWT токена
		// и извлечение ID пользователя из него
		
		// Для данного примера просто устанавливаем ID абитуриента в контексте
		// Примечание: в реальном приложении нужно реализовать проверку JWT
		c.Set("applicant_id", uint(1))
		
		c.Next()
	}
}

// getApplicantID возвращает ID абитуриента из контекста
func getApplicantID(c *gin.Context) uint {
	applicantID, exists := c.Get("applicant_id")
	if !exists {
		return 0
	}
	return applicantID.(uint)
} 