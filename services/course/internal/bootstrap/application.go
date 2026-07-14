package bootstrap

import (
	"github.com/MinhTuan120704/learning-platform/services/course/internal/config"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/storage/postgres"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/storage/redis"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Config *config.Config

	DB *postgres.Client

	Redis *redis.Client

	Router *gin.Engine
}
