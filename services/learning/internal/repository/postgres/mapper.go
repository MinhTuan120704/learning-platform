package postgres

import (
	"time"

	"github.com/MinhTuan120704/learning-platform/services/learning/internal/domain"
	"github.com/google/uuid"
)

type enrollmentRow struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	CourseID   uuid.UUID
	EnrolledAt time.Time
}

func mapEnrollment(row enrollmentRow) *domain.Enrollment {
	return &domain.Enrollment{
		ID:         row.ID,
		UserID:     row.UserID,
		CourseID:   row.CourseID,
		EnrolledAt: row.EnrolledAt,
	}
}
