package postgres

import (
	"context"
	"errors"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repository.PermissionRepository = (*PermissionRepository)(nil)

type PermissionRepository struct {
	db *pgxpool.Pool
}

func NewPermissionRepository(db *pgxpool.Pool) *PermissionRepository {
	return &PermissionRepository{db: db}
}

func (r *PermissionRepository) Create(ctx context.Context, permission *domain.Permission) error {
	const q = `
		INSERT INTO permission (id, code, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, q, permission.ID, permission.Code, permission.Description, permission.CreatedAt, permission.UpdatedAt)
	return err
}

func (r *PermissionRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Permission, error) {
	const q = `
		SELECT id, code, description, created_at, updated_at
		FROM permission
		WHERE id = $1
	`
	var row permissionRow
	err := r.db.QueryRow(ctx, q, id).Scan(&row.ID, &row.Code, &row.Description, &row.CreatedAt, &row.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrPermissionNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapPermission(row), nil
}

func (r *PermissionRepository) FindByCode(ctx context.Context, code string) (*domain.Permission, error) {
	const q = `
		SELECT id, code, description, created_at, updated_at
		FROM permission
		WHERE code = $1
	`
	var row permissionRow
	err := r.db.QueryRow(ctx, q, code).Scan(&row.ID, &row.Code, &row.Description, &row.CreatedAt, &row.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrPermissionNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapPermission(row), nil
}

func (r *PermissionRepository) List(ctx context.Context) ([]domain.Permission, error) {
	const q = `
		SELECT id, code, description, created_at, updated_at
		FROM permission
		ORDER BY code
	`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []domain.Permission
	for rows.Next() {
		var row permissionRow
		if err := rows.Scan(&row.ID, &row.Code, &row.Description, &row.CreatedAt, &row.UpdatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, *mapPermission(row))
	}
	return permissions, rows.Err()
}

func (r *PermissionRepository) Update(ctx context.Context, permission *domain.Permission) error {
	const q = `
		UPDATE permission
		SET code = $1, description = $2, updated_at = $3
		WHERE id = $4
	`
	tag, err := r.db.Exec(ctx, q, permission.Code, permission.Description, permission.UpdatedAt, permission.ID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrPermissionNotFound
	}
	return nil
}

func (r *PermissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `DELETE FROM permission WHERE id = $1`
	tag, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrPermissionNotFound
	}
	return nil
}
