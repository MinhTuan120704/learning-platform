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

var _ repository.CourseRepository = (*CourseRepository)(nil)

type CourseRepository struct {
	db *pgxpool.Pool
}

func NewCourseRepository(db *pgxpool.Pool) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) Create(ctx context.Context, course *domain.Course) error {
	const q = `
		INSERT INTO course (id, category_id, title, slug, description, thumbnail, published, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.Exec(ctx, q,
		course.ID, course.CategoryID, course.Title, course.Slug, course.Description,
		course.Thumbnail, course.Published, course.CreatedAt, course.UpdatedAt,
	)
	return err
}

func (r *CourseRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	const q = `
		SELECT id, category_id, title, slug, description, thumbnail, published, created_at, updated_at, deleted_at
		FROM course
		WHERE id = $1 AND deleted_at IS NULL
	`
	var row courseRow
	err := r.db.QueryRow(ctx, q, id).Scan(
		&row.ID, &row.CategoryID, &row.Title, &row.Slug, &row.Description,
		&row.Thumbnail, &row.Published, &row.CreatedAt, &row.UpdatedAt, &row.DeletedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrCourseNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapCourse(row), nil
}

func (r *CourseRepository) FindBySlug(ctx context.Context, slug string) (*domain.Course, error) {
	const q = `
		SELECT id, category_id, title, slug, description, thumbnail, published, created_at, updated_at, deleted_at
		FROM course
		WHERE slug = $1 AND deleted_at IS NULL
	`
	var row courseRow
	err := r.db.QueryRow(ctx, q, slug).Scan(
		&row.ID, &row.CategoryID, &row.Title, &row.Slug, &row.Description,
		&row.Thumbnail, &row.Published, &row.CreatedAt, &row.UpdatedAt, &row.DeletedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrCourseNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapCourse(row), nil
}

func (r *CourseRepository) List(ctx context.Context) ([]domain.Course, error) {
	const q = `
		SELECT id, category_id, title, slug, description, thumbnail, published, created_at, updated_at, deleted_at
		FROM course
		WHERE deleted_at IS NULL
		ORDER BY created_at DESC
	`
	rows, err := r.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []domain.Course
	for rows.Next() {
		var row courseRow
		if err := rows.Scan(
			&row.ID, &row.CategoryID, &row.Title, &row.Slug, &row.Description,
			&row.Thumbnail, &row.Published, &row.CreatedAt, &row.UpdatedAt, &row.DeletedAt,
		); err != nil {
			return nil, err
		}
		courses = append(courses, *mapCourse(row))
	}
	return courses, rows.Err()
}

func (r *CourseRepository) Update(ctx context.Context, course *domain.Course) error {
	const q = `
		UPDATE course
		SET category_id = $1, title = $2, slug = $3, description = $4,
		    thumbnail = $5, published = $6, updated_at = $7
		WHERE id = $8 AND deleted_at IS NULL
	`
	tag, err := r.db.Exec(ctx, q,
		course.CategoryID, course.Title, course.Slug, course.Description,
		course.Thumbnail, course.Published, course.UpdatedAt, course.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrCourseNotFound
	}
	return nil
}

func (r *CourseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `
		UPDATE course
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	tag, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrCourseNotFound
	}
	return nil
}
