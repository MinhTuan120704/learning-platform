package repository

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/google/uuid"
)

type CourseRepository interface {
	Create(ctx context.Context, course *domain.Course) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Course, error)
	FindBySlug(ctx context.Context, slug string) (*domain.Course, error)
	List(ctx context.Context) ([]domain.Course, error)
	Update(ctx context.Context, course *domain.Course) error
	Delete(ctx context.Context, id uuid.UUID) error
}
