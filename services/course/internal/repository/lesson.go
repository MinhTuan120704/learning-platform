package repository

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/google/uuid"
)

type LessonRepository interface {
	Create(ctx context.Context, lesson *domain.Lesson) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Lesson, error)
	ListBySectionID(ctx context.Context, sectionID uuid.UUID) ([]domain.Lesson, error)
	Update(ctx context.Context, lesson *domain.Lesson) error
	Delete(ctx context.Context, id uuid.UUID) error
}
