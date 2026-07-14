package repository

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/learning/internal/domain"
	"github.com/google/uuid"
)

type EnrollmentRepository interface {
	Create(ctx context.Context, enrollment *domain.Enrollment) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Enrollment, error)
	FindByUser(ctx context.Context, userID uuid.UUID) ([]domain.Enrollment, error)
	Exists(ctx context.Context, userID, courseID uuid.UUID) (bool, error)
	Delete(ctx context.Context, userID, courseID uuid.UUID) error
}
