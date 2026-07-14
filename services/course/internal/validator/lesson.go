package validator

import (
	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/dto"
)

func ValidateCreateLesson(req dto.CreateLessonRequest) error {
	if req.Title == "" {
		return domain.ErrLessonTitleRequired
	}
	if req.Duration < 0 {
		return domain.ErrInvalidPosition
	}
	return nil
}

func ValidateUpdateLesson(req dto.UpdateLessonRequest) error {
	if req.Title != nil && *req.Title == "" {
		return domain.ErrLessonTitleRequired
	}
	if req.Duration != nil && *req.Duration < 0 {
		return domain.ErrInvalidPosition
	}
	return nil
}
