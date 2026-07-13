package http

import (
	handler "github.com/MinhTuan120704/learning-platform/services/identity/internal/http/handler/health"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	healthHandler := handler.NewHealthHandler()

	router.GET("/health", healthHandler.Health)

	return router
}
