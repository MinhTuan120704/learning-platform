package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/MinhTuan120704/learning-platform/services/learning/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/dto"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/service"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/validator"
)

type EnrollmentHandler struct {
	svc *service.EnrollmentService
}

func NewEnrollmentHandler(svc *service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{svc: svc}
}

func (h *EnrollmentHandler) Create(c *gin.Context) {
	userID, err := uuid.Parse(c.GetHeader("X-User-ID"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid user identity"})
		return
	}

	var req dto.CreateEnrollmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := validator.ValidateCreateEnrollment(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	enrollment, err := h.svc.Enroll(c.Request.Context(), userID, req.CourseID)
	if err != nil {
		writeEnrollmentError(c, err)
		return
	}

	c.JSON(http.StatusCreated, dto.EnrollmentResponse{
		CourseID:   enrollment.CourseID,
		EnrolledAt: enrollment.EnrolledAt,
	})
}

func (h *EnrollmentHandler) ListMine(c *gin.Context) {
	userID, err := uuid.Parse(c.GetHeader("X-User-ID"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid user identity"})
		return
	}

	enrollments, err := h.svc.ListMyCourses(c.Request.Context(), userID)
	if err != nil {
		writeEnrollmentError(c, err)
		return
	}

	res := make([]dto.EnrollmentResponse, 0, len(enrollments))
	for _, e := range enrollments {
		res = append(res, dto.EnrollmentResponse{
			CourseID:   e.CourseID,
			EnrolledAt: e.EnrolledAt,
		})
	}
	c.JSON(http.StatusOK, res)
}

func (h *EnrollmentHandler) Delete(c *gin.Context) {
	userID, err := uuid.Parse(c.GetHeader("X-User-ID"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid user identity"})
		return
	}

	courseID, err := uuid.Parse(c.Param("courseId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
		return
	}

	if err := h.svc.Unenroll(c.Request.Context(), userID, courseID); err != nil {
		writeEnrollmentError(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func writeEnrollmentError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrEnrollmentNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrAlreadyEnrolled):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, domain.ErrUserIDRequired), errors.Is(err, domain.ErrCourseIDRequired):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}
}
