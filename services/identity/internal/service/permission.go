package service

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/repository"
	"github.com/google/uuid"
)

type PermissionService struct {
	repo repository.PermissionRepository
}

func NewPermissionService(repo repository.PermissionRepository) *PermissionService {
	return &PermissionService{repo: repo}
}

func (s *PermissionService) GetUserPermissions(ctx context.Context, userID uuid.UUID) ([]string, []string, error) {
	return s.repo.FindByUserID(ctx, userID)
}
