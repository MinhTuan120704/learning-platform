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

var _ repository.LessonRepository = (*LessonRepository)(nil)

type LessonRepository struct {
	db *pgxpool.Pool
}

func NewLessonRepository(db *pgxpool.Pool) *LessonRepository {
	return &LessonRepository{db: db}
}

func (r *LessonRepository) Create(ctx context.Context, lesson *domain.Lesson) error {
	const q = `
		INSERT INTO lesson (id, section_id, title, content, video_url, duration, position, published, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	`
	_, err := r.db.Exec(ctx, q,
		lesson.ID, lesson.SectionID, lesson.Title, lesson.Content, lesson.VideoURL,
		lesson.Duration, lesson.Position, lesson.Published, lesson.CreatedAt, lesson.UpdatedAt,
	)
	return err
}

func (r *LessonRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Lesson, error) {
	const q = `
		SELECT id, section_id, title, content, video_url, duration, position, published, created_at, updated_at
		FROM lesson
		WHERE id = $1
	`
	var row lessonRow
	err := r.db.QueryRow(ctx, q, id).Scan(
		&row.ID, &row.SectionID, &row.Title, &row.Content, &row.VideoURL,
		&row.Duration, &row.Position, &row.Published, &row.CreatedAt, &row.UpdatedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrLessonNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapLesson(row), nil
}

func (r *LessonRepository) ListBySectionID(ctx context.Context, sectionID uuid.UUID) ([]domain.Lesson, error) {
	const q = `
		SELECT id, section_id, title, content, video_url, duration, position, published, created_at, updated_at
		FROM lesson
		WHERE section_id = $1
		ORDER BY position
	`
	rows, err := r.db.Query(ctx, q, sectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []domain.Lesson
	for rows.Next() {
		var row lessonRow
		if err := rows.Scan(
			&row.ID, &row.SectionID, &row.Title, &row.Content, &row.VideoURL,
			&row.Duration, &row.Position, &row.Published, &row.CreatedAt, &row.UpdatedAt,
		); err != nil {
			return nil, err
		}
		lessons = append(lessons, *mapLesson(row))
	}
	return lessons, rows.Err()
}

func (r *LessonRepository) Update(ctx context.Context, lesson *domain.Lesson) error {
	const q = `
		UPDATE lesson
		SET title = $1, content = $2, video_url = $3, duration = $4,
		    position = $5, published = $6, updated_at = $7
		WHERE id = $8
	`
	tag, err := r.db.Exec(ctx, q,
		lesson.Title, lesson.Content, lesson.VideoURL, lesson.Duration,
		lesson.Position, lesson.Published, lesson.UpdatedAt, lesson.ID,
	)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrLessonNotFound
	}
	return nil
}

func (r *LessonRepository) Delete(ctx context.Context, id uuid.UUID) error {
	const q = `DELETE FROM lesson WHERE id = $1`
	tag, err := r.db.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrLessonNotFound
	}
	return nil
}
