package bootstrap

import (
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/config"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/http"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/storage/postgres"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/storage/redis"
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

	redisClient, err := redis.New(cfg.Redis)
	if err != nil {
		return nil, err
	}

	router := http.NewRouter()

	app := &Application{
		Config: cfg,
		DB:     db,
		Redis:  redisClient,
		Router: router,
	}

	return app, nil
}
