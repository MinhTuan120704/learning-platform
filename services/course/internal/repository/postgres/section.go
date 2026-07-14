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

var _ repository.SectionRepository = (*SectionRepository)(nil)

type SectionRepository struct {
	db *pgxpool.Pool
}

func NewSectionRepository(db *pgxpool.Pool) *SectionRepository {
	return &SectionRepository{db: db}
}

func (r *SectionRepository) Create(ctx context.Context, section *domain.Section) error {
	const q = `
		INSERT INTO section (id, course_id, title, description, position, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(ctx, q,
		section.ID, section.CourseID, section.Title, section.Description,
		section.Position, section.CreatedAt, section.UpdatedAt,
	)
	return err
}

func (r *SectionRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Section, error) {
	const q = `
		SELECT id, course_id, title, description, position, created_at, updated_at
		FROM section
		WHERE id = $1
	`
	var row sectionRow
	err := r.db.QueryRow(ctx, q, id).Scan(
		&row.ID, &row.CourseID, &row.Title, &row.Description, &row.Position, &row.CreatedAt, &row.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrSectionNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapSection(row), nil
}

func (r *SectionRepository) ListByCourseID(ctx context.Context, courseID uuid.UUID) ([]domain.Section, error) {
	const q = `
		SELECT id, course_id, title, description, position, created_at, updated_at
		FROM section
		WHERE course_id = $1
		ORDER BY position
	`
	rows, err := r.db.Query(ctx, q, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sections []domain.Section
	for rows.Next() {
		var row sectionRow
		if err := rows.Scan(&row.ID, &row.CourseID, &row.Title, &row.Description, &row.Position, &row.CreatedAt, &row.UpdatedAt); err != nil {
			return nil, err
		}
		sections = append(sections, *mapSection(row))
	}
	return sections, rows.Err()
}

func (r *SectionRepository) Update(ctx context.Context, section *domain.Section) error {
	const q = `
		UPDATE section
		SET title = $1, description = $2, position = $3, updated_at = $4
		WHERE id = $5
	`
	tag, err := r.db.Exec(ctx, q,
		section.Title, section.Description, section.Position, section.UpdatedAt, section.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrSectionNotFound
	}
	return nil
}

func (r *SectionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `DELETE FROM section WHERE id = $1`
	tag, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrSectionNotFound
	}
	return nil
}
