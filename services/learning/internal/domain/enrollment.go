package domain

import (
	"time"

	"github.com/google/uuid"
)

type Enrollment struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	CourseID   uuid.UUID
	EnrolledAt time.Time
}

func NewEnrollment(userID, courseID uuid.UUID) (*Enrollment, error) {
	if userID == uuid.Nil {
		return nil, ErrUserIDRequired
	}
	if courseID == uuid.Nil {
		return nil, ErrCourseIDRequired
	}

	return &Enrollment{
		ID:         uuid.New(),
		UserID:     userID,
		CourseID:   courseID,
		EnrolledAt: time.Now(),
	}, nil
}
