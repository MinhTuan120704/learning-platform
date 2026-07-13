package http

import (
	handler "github.com/MinhTuan120704/learning-platform/services/identity/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	Auth   *handler.AuthHandler
	Health *handler.HealthHandler
}

func NewRouter(deps RouterDeps) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", deps.Health.Check)

	v1 := router.Group("/api/v1")
	{
		authGroup := v1.Group("/auth")
		{
			authGroup.POST("/register", deps.Auth.Register)
			authGroup.POST("/login", deps.Auth.Login)
			authGroup.POST("/refresh", deps.Auth.Refresh)
			authGroup.POST("/logout", deps.Auth.Logout)
		}
	}

	return router
}
