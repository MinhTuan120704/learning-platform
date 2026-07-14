package http

import (
	handler "github.com/MinhTuan120704/learning-platform/services/learning/internal/http/handler"
	"github.com/gin-gonic/gin"
)

type RouterDeps struct {
	Health     *handler.HealthHandler
	Enrollment *handler.EnrollmentHandler
}

func NewRouter(deps RouterDeps) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", deps.Health.Check)

	v1 := router.Group("/api/v1")
	{
		enrollments := v1.Group("/enrollments")
		{
			enrollments.POST("", deps.Enrollment.Create)
			enrollments.GET("/me", deps.Enrollment.ListMine)
			enrollments.DELETE("/:courseId", deps.Enrollment.Delete)
		}
	}

	return router
}
