package repository

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateProfile(ctx context.Context, id uuid.UUID, name string) error
	UpdateEmail(ctx context.Context, id uuid.UUID, email string) error
	UpdatePassword(ctx context.Context, id uuid.UUID, passwordHash string) error
	Delete(ctx context.Context, id uuid.UUID) error
	AssignRole(ctx context.Context, id, roleID uuid.UUID) error
}
