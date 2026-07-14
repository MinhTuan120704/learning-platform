package service

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	users repository.UserRepository
}

func NewUserService(users repository.UserRepository) *UserService {
	return &UserService{users: users}
}

func (s *UserService) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	return s.users.FindByID(ctx, id)
}

func (s *UserService) UpdateProfile(ctx context.Context, id uuid.UUID, name string) (*domain.User, error) {
	if err := s.users.UpdateProfile(ctx, id, name); err != nil {
		return nil, err
	}
	return s.users.FindByID(ctx, id)
}
