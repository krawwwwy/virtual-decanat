package main

import (
	"club-service/internal/handler"
	"club-service/internal/repository"
	"club-service/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	// Initialize database connection
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Failed to initialize db: %s", err.Error())
	}

	// Initialize repositories
	clubRepo := repository.NewClubRepository(db)
	memberRepo := repository.NewMemberRepository(db)

	// Initialize services
	clubService := service.NewClubService(clubRepo)
	memberService := service.NewMemberService(memberRepo)

	// Initialize handlers
	clubHandler := handler.NewClubHandler(clubService)
	memberHandler := handler.NewMemberHandler(memberService)

	// Initialize router
	router := gin.Default()

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Club routes
	clubRoutes := router.Group("/api/clubs")
	{
		clubRoutes.POST("/", clubHandler.CreateClub)
		clubRoutes.GET("/", clubHandler.GetAllClubs)
		clubRoutes.GET("/:id", clubHandler.GetClubByID)
		clubRoutes.PUT("/:id", clubHandler.UpdateClub)
		clubRoutes.DELETE("/:id", clubHandler.DeleteClub)
	}

	// Member routes
	memberRoutes := router.Group("/api/members")
	{
		memberRoutes.POST("/", memberHandler.AddMember)
		memberRoutes.GET("/club/:clubId", memberHandler.GetClubMembers)
		memberRoutes.DELETE("/:id", memberHandler.RemoveMember)
	}

	// Serve frontend
	router.StaticFile("/", "./frontend.html")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %s", err.Error())
	}
} 