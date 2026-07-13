package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	const q = `
		INSERT INTO "user" (id, name, email, password_hash, email_verified_at, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := r.db.Exec(ctx, q, user.ID, user.Name, user.Email, user.PasswordHash, user.EmailVerifiedAt, user.CreatedAt, user.UpdatedAt)

	return err
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	const q = `
		SELECT id, name, email, email_verified_at, created_at, updated_at
		FROM "user" 
		WHERE id = $1 AND deleted_at IS NULL`

	var row userRow
	err := r.db.QueryRow(ctx, q, id).Scan(
		&row.ID, &row.Name, &row.Email, &row.EmailVerifiedAt, &row.CreatedAt, &row.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	roles, err := r.findRolesByUserID(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return mapUser(row, roles), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	const q = `
		SELECT id, name, email, email_verified_at, created_at, updated_at
		FROM "user" 
		WHERE email = $1 AND deleted_at IS NULL`

	var row userRow
	err := r.db.QueryRow(ctx, q, email).Scan(
		&row.ID, &row.Name, &row.Email, &row.EmailVerifiedAt, &row.CreatedAt, &row.UpdatedAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}

	roles, err := r.findRolesByUserID(ctx, row.ID)
	if err != nil {
		return nil, err
	}

	return mapUser(row, roles), nil
}

func (r *UserRepository) UpdateProfile(ctx context.Context, id uuid.UUID, name string) error {
	const q = `UPDATE "user" SET name = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL`
	_, err := r.db.Exec(ctx, q, name, time.Now(), id)
	return err
}

func (r *UserRepository) UpdateEmail(ctx context.Context, id uuid.UUID, email string) error {
	const q = `UPDATE "user" SET email = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL`
	_, err := r.db.Exec(ctx, q, email, time.Now(), id)
	return err
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id uuid.UUID, passwordHash string) error {
	const q = `UPDATE "user" SET password_hash = $1, updated_at = $2 WHERE id = $3 AND deleted_at IS NULL`
	_, err := r.db.Exec(ctx, q, passwordHash, time.Now(), id)
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `
		UPDATE "user"
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	tag, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

func (r *UserRepository) findRolesByUserID(ctx context.Context, userID uuid.UUID) ([]domain.Role, error) {
	const q = `
		SELECT r.id, r.name
		FROM role r
		JOIN user_role ur ON ur.role_id = r.id
		WHERE ur.user_id = $1	
	`

	rows, err := r.db.Query(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		var row roleRow
		if err := rows.Scan(&row.ID, &row.Name); err != nil {
			return nil, err
		}
		roles = append(roles, *mapRole(row, nil))
	}

	return roles, rows.Err()
}
