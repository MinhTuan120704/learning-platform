package validator

import (
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/learning/internal/dto"
	"github.com/google/uuid"
)

func ValidateCreateEnrollment(req dto.CreateEnrollmentRequest) error {
	if req.CourseID == uuid.Nil {
		return domain.ErrCourseIDRequired
	}
	return nil
}
