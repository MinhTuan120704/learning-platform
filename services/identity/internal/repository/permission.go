package repository

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/google/uuid"
)

type PermissionRepository interface {
	Create(ctx context.Context, permission *domain.Permission) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Permission, error)
	FindByUserID(ctx context.Context, userID uuid.UUID) (roles []string, permissions []string, err error)
	FindByCode(ctx context.Context, code string) (*domain.Permission, error)
	List(ctx context.Context) ([]domain.Permission, error)
	Update(ctx context.Context, permission *domain.Permission) error
	Delete(ctx context.Context, id uuid.UUID) error
}
