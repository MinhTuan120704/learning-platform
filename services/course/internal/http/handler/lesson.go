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

type LessonHandler struct {
	svc *service.LessonService
}

func NewLessonHandler(svc *service.LessonService) *LessonHandler {
	return &LessonHandler{svc: svc}
}

func (h *LessonHandler) Create(c *gin.Context) {
	sectionID, err := uuid.Parse(c.Param("sectionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid section id"})
		return
	}
	var req dto.CreateLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateCreateLesson(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lesson, err := h.svc.Create(c.Request.Context(), sectionID, req.Title, req.Content, req.VideoURL, req.Duration)
	if err != nil {
		writeLessonError(c, err)
		return
	}
	c.JSON(http.StatusCreated, toLessonResponse(lesson))
}

func (h *LessonHandler) ListBySection(c *gin.Context) {
	sectionID, err := uuid.Parse(c.Param("sectionId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid section id"})
		return
	}
	lessons, err := h.svc.ListBySection(c.Request.Context(), sectionID)
	if err != nil {
		writeLessonError(c, err)
		return
	}
	res := make([]dto.LessonResponse, 0, len(lessons))
	for _, le := range lessons {
		res = append(res, toLessonResponse(&le))
	}
	c.JSON(http.StatusOK, res)
}

func (h *LessonHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	lesson, err := h.svc.Get(c.Request.Context(), id)
	if err != nil {
		writeLessonError(c, err)
		return
	}
	c.JSON(http.StatusOK, toLessonResponse(lesson))
}

func (h *LessonHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.UpdateLessonRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateUpdateLesson(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lesson, err := h.svc.Update(c.Request.Context(), id, req.Title, req.Content, req.VideoURL, req.Duration)
	if err != nil {
		writeLessonError(c, err)
		return
	}
	c.JSON(http.StatusOK, toLessonResponse(lesson))
}

func (h *LessonHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		writeLessonError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *LessonHandler) Publish(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	lesson, err := h.svc.Publish(c.Request.Context(), id)
	if err != nil {
		writeLessonError(c, err)
		return
	}
	c.JSON(http.StatusOK, toLessonResponse(lesson))
}

func (h *LessonHandler) Unpublish(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	lesson, err := h.svc.Unpublish(c.Request.Context(), id)
	if err != nil {
		writeLessonError(c, err)
		return
	}
	c.JSON(http.StatusOK, toLessonResponse(lesson))
}

func toLessonResponse(le *domain.Lesson) dto.LessonResponse {
	return dto.LessonResponse{
		ID:        le.ID,
		SectionID: le.SectionID,
		Title:     le.Title,
		Content:   le.Content,
		VideoURL:  le.VideoURL,
		Duration:  le.Duration,
		Position:  le.Position,
		Published: le.Published,
	}
}

func writeLessonError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrLessonNotFound), errors.Is(err, domain.ErrSectionNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrLessonTitleRequired):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrLessonAlreadyPublished), errors.Is(err, domain.ErrLessonNotPublished):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
