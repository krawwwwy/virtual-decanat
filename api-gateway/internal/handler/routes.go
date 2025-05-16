package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/krawwwwy/virtual-decanat/api-gateway/internal/middleware"
	"go.uber.org/zap"
)

// SetupRouter настраивает маршруты API Gateway
func SetupRouter(logger *zap.Logger) *gin.Engine {
	router := gin.New()
	
	// Применяем middleware
	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(middleware.TracingMiddleware("api-gateway"))

	// Создаем обработчик прокси
	proxyHandler := NewProxyHandler(logger)

	// Публичные маршруты
	public := router.Group("/api")
	{
		// Маршруты аутентификации
		auth := public.Group("/auth")
		{
			auth.POST("/login", proxyHandler.ProxyRequest("auth_service"))
			auth.POST("/register", proxyHandler.ProxyRequest("auth_service"))
			auth.POST("/refresh", proxyHandler.ProxyRequest("auth_service"))
		}

		// Маршруты для абитуриентов
		applicant := public.Group("/applicant")
		{
			applicant.POST("/apply", proxyHandler.ProxyRequest("applicant_service"))
			applicant.GET("/status/:id", proxyHandler.ProxyRequest("applicant_service"))
		}

		// Публичное расписание
		public.GET("/schedule/public", proxyHandler.ProxyRequest("schedule_service"))
	}

	// Защищенные маршруты (требуют аутентификации)
	protected := router.Group("/api")
	protected.Use(middleware.JWTAuthMiddleware())
	{
		// Профиль пользователя
		protected.GET("/profile", proxyHandler.ProxyRequest("auth_service"))
		protected.PUT("/profile", proxyHandler.ProxyRequest("auth_service"))

		// Расписание
		schedule := protected.Group("/schedule")
		{
			schedule.GET("/", proxyHandler.ProxyRequest("schedule_service"))
			
			// Маршруты для преподавателей и сотрудников деканата
			scheduleAdmin := schedule.Group("/")
			scheduleAdmin.Use(middleware.RoleAuthMiddleware("teacher", "dean_office", "admin"))
			{
				scheduleAdmin.POST("/", proxyHandler.ProxyRequest("schedule_service"))
				scheduleAdmin.PUT("/:id", proxyHandler.ProxyRequest("schedule_service"))
				scheduleAdmin.DELETE("/:id", proxyHandler.ProxyRequest("schedule_service"))
			}
		}

		// Успеваемость
		performance := protected.Group("/performance")
		{
			// Маршруты для студентов
			performance.GET("/grades", proxyHandler.ProxyRequest("performance_service"))
			performance.GET("/attendance", proxyHandler.ProxyRequest("performance_service"))
			performance.GET("/debts", proxyHandler.ProxyRequest("performance_service"))
			performance.GET("/rating", proxyHandler.ProxyRequest("performance_service"))
			
			// Маршруты для преподавателей
			gradesAdmin := performance.Group("/")
			gradesAdmin.Use(middleware.RoleAuthMiddleware("teacher", "dean_office", "admin"))
			{
				gradesAdmin.POST("/grades", proxyHandler.ProxyRequest("performance_service"))
				gradesAdmin.PUT("/grades/:id", proxyHandler.ProxyRequest("performance_service"))
				gradesAdmin.POST("/attendance", proxyHandler.ProxyRequest("performance_service"))
				gradesAdmin.PUT("/attendance/:id", proxyHandler.ProxyRequest("performance_service"))
			}
			
			// Маршруты для сотрудников деканата
			deanOffice := performance.Group("/")
			deanOffice.Use(middleware.RoleAuthMiddleware("dean_office", "admin"))
			{
				deanOffice.POST("/debts", proxyHandler.ProxyRequest("performance_service"))
				deanOffice.PUT("/debts/:id", proxyHandler.ProxyRequest("performance_service"))
				deanOffice.DELETE("/debts/:id", proxyHandler.ProxyRequest("performance_service"))
			}
		}

		// Студенческие объединения
		clubs := protected.Group("/clubs")
		{
			clubs.GET("/", proxyHandler.ProxyRequest("club_service"))
			clubs.GET("/:id", proxyHandler.ProxyRequest("club_service"))
			clubs.POST("/apply/:id", proxyHandler.ProxyRequest("club_service"))
			clubs.GET("/applications", proxyHandler.ProxyRequest("club_service"))
			
			// Маршруты для администраторов клубов
			clubsAdmin := clubs.Group("/")
			clubsAdmin.Use(middleware.RoleAuthMiddleware("dean_office", "admin"))
			{
				clubsAdmin.POST("/", proxyHandler.ProxyRequest("club_service"))
				clubsAdmin.PUT("/:id", proxyHandler.ProxyRequest("club_service"))
				clubsAdmin.DELETE("/:id", proxyHandler.ProxyRequest("club_service"))
				clubsAdmin.PUT("/applications/:id", proxyHandler.ProxyRequest("club_service"))
			}
		}

		// Социальная поддержка
		support := protected.Group("/support")
		{
			support.GET("/types", proxyHandler.ProxyRequest("support_service"))
			support.POST("/apply", proxyHandler.ProxyRequest("support_service"))
			support.GET("/applications", proxyHandler.ProxyRequest("support_service"))
			
			// Маршруты для сотрудников деканата
			supportAdmin := support.Group("/")
			supportAdmin.Use(middleware.RoleAuthMiddleware("dean_office", "admin"))
			{
				supportAdmin.POST("/types", proxyHandler.ProxyRequest("support_service"))
				supportAdmin.PUT("/types/:id", proxyHandler.ProxyRequest("support_service"))
				supportAdmin.DELETE("/types/:id", proxyHandler.ProxyRequest("support_service"))
				supportAdmin.PUT("/applications/:id", proxyHandler.ProxyRequest("support_service"))
			}
		}
	}

	return router
} 