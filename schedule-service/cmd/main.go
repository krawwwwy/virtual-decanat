package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"context"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/model"
	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/repository"
	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/service"
	"github.com/krawwwwy/virtual-decanat/schedule-service/internal/handler"
)

func main() {
	if err := initConfig(); err != nil {
		panic(fmt.Errorf("fatal config: %w", err))
	}
	db, err := initDB()
	if err != nil {
		panic(fmt.Errorf("db: %w", err))
	}

	db.AutoMigrate(&model.Group{}, &model.Subject{}, &model.Schedule{})

	repo := repository.NewScheduleRepository(db)
	svc := service.NewScheduleService(repo)
	h := handler.NewScheduleHandler(svc)

	r := gin.New()
	r.Use(gin.Recovery())

	r.StaticFile("/", "./frontend.html")
	r.StaticFile("/test.html", "./frontend.html")

	api := r.Group("/api/v1")
	{
		schedule := api.Group("/schedule")
		{
			schedule.POST("/", h.CreateSchedule)
			schedule.PUT("/:id", h.UpdateSchedule)
			schedule.DELETE("/:id", h.DeleteSchedule)
			schedule.GET("/:id", h.GetScheduleByID)
			schedule.GET("/teacher/:teacher_id", h.ListByTeacher)
			schedule.GET("/group/:group_id", h.ListByGroup)
		}
		group := api.Group("/group")
		{
			group.POST("/", h.CreateGroup)
			group.PUT("/:id", h.UpdateGroup)
			group.DELETE("/:id", h.DeleteGroup)
			group.GET("/:id", h.GetGroupByID)
			group.GET("/", h.ListGroups)
		}
		subject := api.Group("/subject")
		{
			subject.POST("/", h.CreateSubject)
			subject.PUT("/:id", h.UpdateSubject)
			subject.DELETE("/:id", h.DeleteSubject)
			subject.GET("/:id", h.GetSubjectByID)
			subject.GET("/", h.ListSubjects)
		}
	}

	port := viper.GetString("server.port")
	if port == "" {
		port = "8082"
	}

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("listen: %w", err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func initDB() (*gorm.DB, error) {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	sslmode := viper.GetString("database.sslmode")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
} 