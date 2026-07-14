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

type SectionHandler struct {
	svc *service.SectionService
}

func NewSectionHandler(svc *service.SectionService) *SectionHandler {
	return &SectionHandler{svc: svc}
}

func (h *SectionHandler) Create(c *gin.Context) {
	courseID, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
		return
	}
	var req dto.CreateSectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateCreateSection(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	section, err := h.svc.Create(c.Request.Context(), courseID, req.Title, req.Description)
	if err != nil {
		writeSectionError(c, err)
		return
	}
	c.JSON(http.StatusCreated, toSectionResponse(section))
}

func (h *SectionHandler) ListByCourse(c *gin.Context) {
	courseID, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
		return
	}
	sections, err := h.svc.ListByCourse(c.Request.Context(), courseID)
	if err != nil {
		writeSectionError(c, err)
		return
	}
	res := make([]dto.SectionResponse, 0, len(sections))
	for _, se := range sections {
		res = append(res, toSectionResponse(&se))
	}
	c.JSON(http.StatusOK, res)
}

func (h *SectionHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req dto.UpdateSectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateUpdateSection(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	section, err := h.svc.Update(c.Request.Context(), id, req.Title, req.Description)
	if err != nil {
		writeSectionError(c, err)
		return
	}
	c.JSON(http.StatusOK, toSectionResponse(section))
}

func (h *SectionHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		writeSectionError(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func toSectionResponse(se *domain.Section) dto.SectionResponse {
	return dto.SectionResponse{
		ID:          se.ID,
		CourseID:    se.CourseID,
		Title:       se.Title,
		Description: se.Description,
		Position:    se.Position,
	}
}

func writeSectionError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrSectionNotFound), errors.Is(err, domain.ErrCourseNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrSectionTitleRequired):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
