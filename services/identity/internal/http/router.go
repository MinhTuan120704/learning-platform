package http

import (
	handler "github.com/MinhTuan120704/learning-platform/services/identity/internal/http/handler"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/http/middleware"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/token"
	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	Auth       *handler.AuthHandler
	User       *handler.UserHandler
	Permission *handler.PermissionHandler
	Health     *handler.HealthHandler
}

func NewRouter(deps RouterDeps, internalAPIKey string, jwtSvc *token.JWTService) *gin.Engine {
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

		users := v1.Group("/users")
		users.Use(middleware.Authenticate(jwtSvc))
		{
			users.GET("/me", deps.User.Me)
			users.PATCH("/me", deps.User.UpdateMe)
		}
	}

	internal := router.Group("/internal")
	internal.Use(middleware.RequireServiceAuth(internalAPIKey))
	{
		internal.GET("/permissions", deps.Permission.GetUserPermissions)
	}

	return router
}
