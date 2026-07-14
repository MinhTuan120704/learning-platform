package service

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/repository"
	"github.com/google/uuid"
)

type SectionService struct {
	sections repository.SectionRepository
	courses  repository.CourseRepository
}

func NewSectionService(sections repository.SectionRepository, courses repository.CourseRepository) *SectionService {
	return &SectionService{sections: sections, courses: courses}
}

func (s *SectionService) Create(ctx context.Context, courseID uuid.UUID, title, description string) (*domain.Section, error) {
	if _, err := s.courses.FindByID(ctx, courseID); err != nil {
		return nil, err
	}

	existing, err := s.sections.ListByCourseID(ctx, courseID)
	if err != nil {
		return nil, err
	}
	position := len(existing) + 1

	section, err := domain.NewSection(courseID, title, description, position)
	if err != nil {
		return nil, err
	}
	if err := s.sections.Create(ctx, section); err != nil {
		return nil, err
	}
	return section, nil
}

func (s *SectionService) Get(ctx context.Context, id uuid.UUID) (*domain.Section, error) {
	return s.sections.FindByID(ctx, id)
}

func (s *SectionService) ListByCourse(ctx context.Context, courseID uuid.UUID) ([]domain.Section, error) {
	return s.sections.ListByCourseID(ctx, courseID)
}

func (s *SectionService) Update(ctx context.Context, id uuid.UUID, title, description *string) (*domain.Section, error) {
	section, err := s.sections.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if title != nil {
		if err := section.Rename(*title); err != nil {
			return nil, err
		}
	}
	if description != nil {
		section.UpdateDescription(*description)
	}
	if err := s.sections.Update(ctx, section); err != nil {
		return nil, err
	}
	return section, nil
}

func (s *SectionService) Move(ctx context.Context, id uuid.UUID, position int) (*domain.Section, error) {
	section, err := s.sections.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	section.Move(position)
	if err := s.sections.Update(ctx, section); err != nil {
		return nil, err
	}
	return section, nil
}

func (s *SectionService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.sections.Delete(ctx, id)
}
