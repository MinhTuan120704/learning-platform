package repository

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/google/uuid"
)

type SectionRepository interface {
	Create(ctx context.Context, section *domain.Section) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Section, error)
	ListByCourseID(ctx context.Context, courseID uuid.UUID) ([]domain.Section, error)
	Update(ctx context.Context, section *domain.Section) error
	Delete(ctx context.Context, id uuid.UUID) error
}
