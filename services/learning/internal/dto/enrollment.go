package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateEnrollmentRequest struct {
	CourseID uuid.UUID `json:"course_id" binding:"required"`
}

type EnrollmentResponse struct {
	CourseID   uuid.UUID `json:"course_id"`
	EnrolledAt time.Time `json:"enrolled_at"`
}
