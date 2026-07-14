package validator

import (
	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/dto"
)

func ValidateCreateSection(req dto.CreateSectionRequest) error {
	if req.Title == "" {
		return domain.ErrSectionTitleRequired
	}
	return nil
}

func ValidateUpdateSection(req dto.UpdateSectionRequest) error {
	if req.Title != nil && *req.Title == "" {
		return domain.ErrSectionTitleRequired
	}
	return nil
}
