package bootstrap

import (
	"net/http"
	"time"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/config"
	httpserver "github.com/MinhTuan120704/learning-platform/services/course/internal/http"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/http/handler"
	repopg "github.com/MinhTuan120704/learning-platform/services/course/internal/repository/postgres"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/service"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/storage/postgres"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/storage/redis"
)

func New() (*Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	db, err := postgres.New(cfg.DB)
	if err != nil {
		return nil, err
	}

	redisClient, err := redis.New(cfg.Redis.URL)
	if err != nil {
		return nil, err
	}

	// Repository
	categoryRepo := repopg.NewCategoryRepository(db.Pool)
	courseRepo := repopg.NewCourseRepository(db.Pool)
	sectionRepo := repopg.NewSectionRepository(db.Pool)
	lessonRepo := repopg.NewLessonRepository(db.Pool)

	// Service
	categorySvc := service.NewCategoryService(categoryRepo)
	courseSvc := service.NewCourseService(courseRepo, categoryRepo)
	sectionSvc := service.NewSectionService(sectionRepo, courseRepo)
	lessonSvc := service.NewLessonService(lessonRepo, sectionRepo)

	// Handler
	categoryHandler := handler.NewCategoryHandler(categorySvc)
	courseHandler := handler.NewCourseHandler(courseSvc)
	sectionHandler := handler.NewSectionHandler(sectionSvc)
	lessonHandler := handler.NewLessonHandler(lessonSvc)
	healthHandler := handler.NewHealthHandler()

	httpClient := &http.Client{Timeout: 3 * time.Second}

	router := httpserver.NewRouter(httpserver.RouterDeps{
		Category: categoryHandler,
		Course:   courseHandler,
		Section:  sectionHandler,
		Lesson:   lessonHandler,
		Health:   healthHandler,
	}, redisClient.Client, cfg.HTTP.IdentityServiceUrl, cfg.HTTP.InternalApiKey, httpClient)

	app := &Application{
		Config: cfg,
		DB:     db,
		Redis:  redisClient,
		Router: router,
	}

	return app, nil
}

func (a *Application) Close() {
	if a.DB != nil {
		a.DB.Pool.Close()
	}
	if a.Redis != nil {
		_ = a.Redis.Close()
	}
}
