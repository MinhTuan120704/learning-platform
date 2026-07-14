package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/dto"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/service"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/validator"
)

type CourseHandler struct {
	svc *service.CourseService
}

func NewCourseHandler(svc *service.CourseService) *CourseHandler {
	return &CourseHandler{svc: svc}
}

func (h *CourseHandler) Create(c *gin.Context) {
	var req dto.CreateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateCreateCourse(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.svc.Create(c.Request.Context(), req.CategoryID, req.Title, req.Slug, req.Description)
	if err != nil {
		writeCourseError(c, err)
		return
	}

	c.JSON(http.StatusCreated, toCourseResponse(course))
}

func (h *CourseHandler) List(c *gin.Context) {
	courses, err := h.svc.List(c.Request.Context())
	if err != nil {
		writeCourseError(c, err)
		return
	}

	res := make([]dto.CourseResponse, 0, len(courses))
	for _, co := range courses {
		res = append(res, toCourseResponse(&co))
	}
	c.JSON(http.StatusOK, res)
}

func (h *CourseHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	course, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		writeCourseError(c, err)
		return
	}

	c.JSON(http.StatusOK, toCourseResponse(course))
}

func (h *CourseHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateCourseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateUpdateCourse(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := h.svc.Update(c.Request.Context(), id, req.Title, req.Description)
	if err != nil {
		writeCourseError(c, err)
		return
	}

	c.JSON(http.StatusOK, toCourseResponse(course))
}

func (h *CourseHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		writeCourseError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func toCourseResponse(co *domain.Course) dto.CourseResponse {
	return dto.CourseResponse{
		ID:          co.ID,
		CategoryID:  co.CategoryID,
		Title:       co.Title,
		Slug:        co.Slug,
		Description: co.Description,
		Thumbnail:   co.Thumbnail,
		Published:   co.Published,
	}
}

func writeCourseError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrCourseNotFound), errors.Is(err, domain.ErrCategoryNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrCourseSlugAlreadyExists):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrCourseTitleRequired), errors.Is(err, domain.ErrCourseSlugRequired):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrCourseAlreadyPublished), errors.Is(err, domain.ErrCourseNotPublished):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
