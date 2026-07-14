package bootstrap

import (
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/config"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/storage/postgres"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/storage/redis"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Config *config.Config

	DB *postgres.Client

	Redis *redis.Client

	Router *gin.Engine
}
