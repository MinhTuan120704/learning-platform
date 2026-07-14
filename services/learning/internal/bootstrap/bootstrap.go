package bootstrap

import (
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/config"
	httpserver "github.com/MinhTuan120704/learning-platform/services/learning/internal/http"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/http/handler"
	repopg "github.com/MinhTuan120704/learning-platform/services/learning/internal/repository/postgres"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/service"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/storage/postgres"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/storage/redis"
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
	enrollmentRepo := repopg.NewEnrollmentRepository(db.Pool)

	// Service
	enrollmentSvc := service.NewEnrollmentService(enrollmentRepo)

	// Handler
	enrollmentHandler := handler.NewEnrollmentHandler(enrollmentSvc)
	healthHandler := handler.NewHealthHandler()

	router := httpserver.NewRouter(httpserver.RouterDeps{
		Enrollment: enrollmentHandler,
		Health:     healthHandler,
	})

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
