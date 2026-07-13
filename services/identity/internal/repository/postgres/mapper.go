package postgres

import (
	"time"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/google/uuid"
)

type userRow struct {
	ID              uuid.UUID
	Name            string
	Email           string
	PasswordHash    string
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func mapUser(row userRow, roles []domain.Role) *domain.User {
	return &domain.User{
		ID:              row.ID,
		Name:            row.Name,
		Email:           row.Email,
		PasswordHash:    row.PasswordHash,
		EmailVerifiedAt: row.EmailVerifiedAt,
		CreatedAt:       row.CreatedAt,
		UpdatedAt:       row.UpdatedAt,
		DeletedAt:       row.DeletedAt,
		Roles:           roles,
	}
}

type roleRow struct {
	ID          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func mapRole(row roleRow, permissions []domain.Permission) *domain.Role {
	return &domain.Role{
		ID:          row.ID,
		Name:        row.Name,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		Permissions: permissions,
	}
}

type permissionRow struct {
	ID          uuid.UUID
	Code        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func mapPermission(row permissionRow) *domain.Permission {
	return &domain.Permission{
		ID:          row.ID,
		Code:        row.Code,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}
