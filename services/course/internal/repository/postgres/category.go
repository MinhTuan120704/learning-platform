package postgres

import (
	"context"
	"errors"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repository.CategoryRepository = (*CategoryRepository)(nil)

type CategoryRepository struct {
	db *pgxpool.Pool
}

func NewCategoryRepository(db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) Create(ctx context.Context, category *domain.Category) error {
	const q = `
		INSERT INTO category (id, name, slug, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(ctx, q,
		category.ID, category.Name, category.Slug, category.Description,
		category.CreatedAt, category.UpdatedAt,
	)
	return err
}

func (r *CategoryRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	const q = `
		SELECT id, name, slug, description, created_at, updated_at
		FROM category
		WHERE id = $1
	`
	var row categoryRow
	err := r.db.QueryRow(ctx, q, id).Scan(
		&row.ID, &row.Name, &row.Slug, &row.Description, &row.CreatedAt, &row.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapCategory(row), nil
}

func (r *CategoryRepository) FindBySlug(ctx context.Context, slug string) (*domain.Category, error) {
	const q = `
		SELECT id, name, slug, description, created_at, updated_at
		FROM category
		WHERE slug = $1
	`
	var row categoryRow
	err := r.db.QueryRow(ctx, q, slug).Scan(
		&row.ID, &row.Name, &row.Slug, &row.Description, &row.CreatedAt, &row.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapCategory(row), nil
}

func (r *CategoryRepository) List(ctx context.Context) ([]domain.Category, error) {
	const q = `
		SELECT id, name, slug, description, created_at, updated_at
		FROM category
		ORDER BY name
	`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		var row categoryRow
		if err := rows.Scan(&row.ID, &row.Name, &row.Slug, &row.Description, &row.CreatedAt, &row.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, *mapCategory(row))
	}
	return categories, rows.Err()
}

func (r *CategoryRepository) Update(ctx context.Context, category *domain.Category) error {
	const q = `
		UPDATE category
		SET name = $1, slug = $2, description = $3, updated_at = $4
		WHERE id = $5
	`
	tag, err := r.db.Exec(ctx, q,
		category.Name, category.Slug, category.Description, category.UpdatedAt, category.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrCategoryNotFound
	}
	return nil
}

func (r *CategoryRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `DELETE FROM category WHERE id = $1`
	tag, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrCategoryNotFound
	}
	return nil
}
