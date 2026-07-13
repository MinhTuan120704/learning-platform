package repository

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/google/uuid"
)

type PermissionRepository interface {
	FindByID(ctx context.Context, id uuid.UUID) (*domain.Permission, error)
	FindByCode(ctx context.Context, code string) (*domain.Permission, error)
	List(ctx context.Context) ([]domain.Permission, error)
}
