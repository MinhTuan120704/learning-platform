package bootstrap

import (
	"time"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/config"
	httpserver "github.com/MinhTuan120704/learning-platform/services/identity/internal/http"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/http/handler"
	repopg "github.com/MinhTuan120704/learning-platform/services/identity/internal/repository/postgres"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/security"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/service"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/storage/postgres"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/storage/redis"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/token"
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
	userRepo := repopg.NewUserRepository(db.Pool)
	roleRepo := repopg.NewRoleRepository(db.Pool)
	permissionRepo := repopg.NewPermissionRepository(db.Pool)

	// Security
	passwordSvc := security.NewPasswordService()

	accessTTL := time.Duration(cfg.JWT.AccessExpireMinutes) * time.Minute
	refreshTTL := time.Duration(cfg.JWT.RefreshExpireDays) * 24 * time.Hour

	jwtSvc := token.NewJWTService(cfg.JWT.Secret, accessTTL, cfg.JWT.Issuer)
	refreshSvc := token.NewRefreshTokenService(redisClient.Client, refreshTTL)

	// Service
	authSvc := service.NewAuthService(userRepo, roleRepo, passwordSvc, jwtSvc, refreshSvc)
	permissionSvc := service.NewPermissionService(permissionRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authSvc)
	permissionHandler := handler.NewPermissionHandler(permissionSvc)
	healthHandler := handler.NewHealthHandler()

	router := httpserver.NewRouter(httpserver.RouterDeps{
		Auth:       authHandler,
		Permission: permissionHandler,
		Health:     healthHandler,
	}, cfg.HTTP.InternalApiKey)

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
