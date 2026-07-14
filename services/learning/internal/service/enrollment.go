package service

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/learning/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/repository"
	"github.com/google/uuid"
)

type EnrollmentService struct {
	enrollments repository.EnrollmentRepository
}

func NewEnrollmentService(enrollments repository.EnrollmentRepository) *EnrollmentService {
	return &EnrollmentService{enrollments: enrollments}
}

func (s *EnrollmentService) Enroll(ctx context.Context, userID, courseID uuid.UUID) (*domain.Enrollment, error) {
	exists, err := s.enrollments.Exists(ctx, userID, courseID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, domain.ErrAlreadyEnrolled
	}

	enrollment, err := domain.NewEnrollment(userID, courseID)
	if err != nil {
		return nil, err
	}

	if err := s.enrollments.Create(ctx, enrollment); err != nil {
		return nil, err
	}
	return enrollment, nil
}

func (s *EnrollmentService) ListMyCourses(ctx context.Context, userID uuid.UUID) ([]domain.Enrollment, error) {
	return s.enrollments.FindByUser(ctx, userID)
}

func (s *EnrollmentService) Unenroll(ctx context.Context, userID, courseID uuid.UUID) error {
	return s.enrollments.Delete(ctx, userID, courseID)
}
