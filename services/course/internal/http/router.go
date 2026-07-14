package http

import (
	"net/http"
	"time"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/cache"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/client"
	handler "github.com/MinhTuan120704/learning-platform/services/course/internal/http/handler"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/http/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RouterDeps struct {
	Health   *handler.HealthHandler
	Category *handler.CategoryHandler
	Course   *handler.CourseHandler
	Section  *handler.SectionHandler
	Lesson   *handler.LessonHandler
}

func NewRouter(deps RouterDeps, redisClient *redis.Client, identityServiceURL, internalAPIKey string, httpClient *http.Client) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	permCache := cache.NewPermissionCache(redisClient, 2*time.Minute)
	identityClient := client.NewIdentityClient(identityServiceURL, internalAPIKey, httpClient)
	requirePerm := func(perm string) gin.HandlerFunc {
		return middleware.RequirePermission(permCache, identityClient, perm)
	}

	router.GET("/health", deps.Health.Check)

	v1 := router.Group("/api/v1")
	{
		categories := v1.Group("/categories")
		{
			categories.GET("", deps.Category.List)
			categories.GET("/:id", deps.Category.Get)
			categories.POST("", requirePerm("category.create"), deps.Category.Create)
			categories.PATCH("/:id", requirePerm("category.manage"), deps.Category.Update)
			categories.DELETE("/:id", requirePerm("category.manage"), deps.Category.Delete)
		}

		courses := v1.Group("/courses")
		{
			courses.GET("", deps.Course.List)

			courses.GET("/:courseId", deps.Course.Get)
			courses.POST("", requirePerm("course.create"), deps.Course.Create)
			courses.PATCH("/:courseId", requirePerm("course.manage"), deps.Course.Update)
			courses.DELETE("/:courseId", requirePerm("course.manage"), deps.Course.Delete)

			courses.POST("/:courseId/sections", requirePerm("section.create"), deps.Section.Create)
			courses.GET("/:courseId/sections", deps.Section.ListByCourse)
		}

		sections := v1.Group("/sections")
		{
			sections.PATCH("/:id", requirePerm("section.manage"), deps.Section.Update)
			sections.DELETE("/:id", requirePerm("section.manage"), deps.Section.Delete)

			sections.POST("/:sectionId/lessons", requirePerm("lesson.create"), deps.Lesson.Create)
			sections.GET("/:sectionId/lessons", deps.Lesson.ListBySection)
		}

		lessons := v1.Group("/lessons")
		{
			lessons.GET("/:id", deps.Lesson.Get)
			lessons.PATCH("/:id", requirePerm("lesson.manage"), deps.Lesson.Update)
			lessons.DELETE("/:id", requirePerm("lesson.manage"), deps.Lesson.Delete)
		}
	}

	return router
}
