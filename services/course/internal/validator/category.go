package validator

import (
	"errors"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/dto"
)

var (
	ErrCategoryNameRequired = errors.New("name is required")
	ErrCategorySlugRequired = errors.New("slug is required")
)

func ValidateCreateCategory(req dto.CreateCategoryRequest) error {
	if req.Name == "" {
		return ErrCategoryNameRequired
	}
	if req.Slug == "" {
		return ErrCategorySlugRequired
	}
	return nil
}

func ValidateUpdateCategory(req dto.UpdateCategoryRequest) error {
	if req.Name != nil && *req.Name == "" {
		return ErrCategoryNameRequired
	}
	if req.Slug != nil && *req.Slug == "" {
		return ErrCategorySlugRequired
	}
	return nil
}
