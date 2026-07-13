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

var _ repository.RoleRepository = (*RoleRepository)(nil)

type RoleRepository struct {
	db *pgxpool.Pool
}

func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Create(ctx context.Context, role *domain.Role) error {
	const q = `
		INSERT INTO role (id, name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx, q, role.ID, role.Name, role.Description, role.CreatedAt, role.UpdatedAt)
	return err
}

func (r *RoleRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	const q = `
		SELECT id, name, description, created_at, updated_at
		FROM role
		WHERE id = $1
	`
	var row roleRow
	err := r.db.QueryRow(ctx, q, id).Scan(&row.ID, &row.Name, &row.Description, &row.CreatedAt, &row.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrRoleNotFound
	}
	if err != nil {
		return nil, err
	}

	permissions, err := r.findPermissionsByRoleID(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return mapRole(row, permissions), nil
}

func (r *RoleRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	const q = `
		SELECT id, name, description, created_at, updated_at
		FROM role
		WHERE name = $1
	`
	var row roleRow
	err := r.db.QueryRow(ctx, q, name).Scan(&row.ID, &row.Name, &row.Description, &row.CreatedAt, &row.UpdatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrRoleNotFound
	}
	if err != nil {
		return nil, err
	}

	permissions, err := r.findPermissionsByRoleID(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return mapRole(row, permissions), nil
}

func (r *RoleRepository) List(ctx context.Context) ([]domain.Role, error) {
	const q = `
		SELECT id, name, description, created_at, updated_at
		FROM role
		ORDER BY name
	`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		var row roleRow
		if err := rows.Scan(&row.ID, &row.Name, &row.Description, &row.CreatedAt, &row.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, *mapRole(row, nil))
	}
	return roles, rows.Err()
}

func (r *RoleRepository) Update(ctx context.Context, role *domain.Role) error {
	const q = `
		UPDATE role
		SET name = $1, description = $2, updated_at = $3
		WHERE id = $4
	`
	tag, err := r.db.Exec(ctx, q, role.Name, role.Description, role.UpdatedAt, role.ID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrRoleNotFound
	}
	return nil
}

func (r *RoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `DELETE FROM role WHERE id = $1`
	tag, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrRoleNotFound
	}
	return nil
}

func (r *RoleRepository) findPermissionsByRoleID(ctx context.Context, roleID uuid.UUID) ([]domain.Permission, error) {
	const q = `
		SELECT p.id, p.code, p.description, p.created_at, p.updated_at
		FROM permission p
		JOIN role_permission rp ON rp.permission_id = p.id
		WHERE rp.role_id = $1
	`
	rows, err := r.db.Query(ctx, q, roleID)
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
