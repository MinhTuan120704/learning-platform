package validator

import (
	"errors"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/dto"
	"github.com/google/uuid"
)

var (
	ErrCourseTitleRequired    = errors.New("title is required")
	ErrCourseSlugRequired     = errors.New("slug is required")
	ErrCourseCategoryRequired = errors.New("category_id is required")
)

func ValidateCreateCourse(req dto.CreateCourseRequest) error {
	if req.Title == "" {
		return ErrCourseTitleRequired
	}
	if req.Slug == "" {
		return ErrCourseSlugRequired
	}
	if req.CategoryID == uuid.Nil {
		return ErrCourseCategoryRequired
	}
	return nil
}

func ValidateUpdateCourse(req dto.UpdateCourseRequest) error {
	if req.Title != nil && *req.Title == "" {
		return ErrCourseTitleRequired
	}
	return nil
}
