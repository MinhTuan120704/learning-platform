package postgres

import (
	"context"
	"errors"

	"github.com/MinhTuan120704/learning-platform/services/learning/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/repository"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repository.EnrollmentRepository = (*EnrollmentRepository)(nil)

type EnrollmentRepository struct {
	db *pgxpool.Pool
}

func NewEnrollmentRepository(db *pgxpool.Pool) *EnrollmentRepository {
	return &EnrollmentRepository{db: db}
}

func (r *EnrollmentRepository) Create(ctx context.Context, e *domain.Enrollment) error {
	const q = `
		INSERT INTO enrollment (id, user_id, course_id, enrolled_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, q, e.ID, e.UserID, e.CourseID, e.EnrolledAt)
	return err
}

func (r *EnrollmentRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.Enrollment, error) {
	const q = `
		SELECT id, user_id, course_id, enrolled_at
		FROM enrollment
		WHERE id = $1
	`
	var row enrollmentRow
	err := r.db.QueryRow(ctx, q, id).Scan(&row.ID, &row.UserID, &row.CourseID, &row.EnrolledAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, domain.ErrEnrollmentNotFound
	}
	if err != nil {
		return nil, err
	}
	return mapEnrollment(row), nil
}

func (r *EnrollmentRepository) FindByUser(ctx context.Context, userID uuid.UUID) ([]domain.Enrollment, error) {
	const q = `
		SELECT id, user_id, course_id, enrolled_at
		FROM enrollment
		WHERE user_id = $1
		ORDER BY enrolled_at DESC
	`
	rows, err := r.db.Query(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Enrollment
	for rows.Next() {
		var row enrollmentRow
		if err := rows.Scan(&row.ID, &row.UserID, &row.CourseID, &row.EnrolledAt); err != nil {
			return nil, err
		}
		result = append(result, *mapEnrollment(row))
	}
	return result, rows.Err()
}

func (r *EnrollmentRepository) Exists(ctx context.Context, userID, courseID uuid.UUID) (bool, error) {
	const q = `SELECT EXISTS(SELECT 1 FROM enrollment WHERE user_id = $1 AND course_id = $2)`
	var exists bool
	err := r.db.QueryRow(ctx, q, userID, courseID).Scan(&exists)
	return exists, err
}

func (r *EnrollmentRepository) Delete(ctx context.Context, userID, courseID uuid.UUID) error {
	const q = `DELETE FROM enrollment WHERE user_id = $1 AND course_id = $2`
	tag, err := r.db.Exec(ctx, q, userID, courseID)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return domain.ErrEnrollmentNotFound
	}
	return nil
}
