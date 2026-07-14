package service

import (
	"context"
	"errors"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/repository"
	"github.com/google/uuid"
)

type CourseService struct {
	courses    repository.CourseRepository
	categories repository.CategoryRepository
}

func NewCourseService(courses repository.CourseRepository, categories repository.CategoryRepository) *CourseService {
	return &CourseService{courses: courses, categories: categories}
}

func (s *CourseService) Create(ctx context.Context, categoryID uuid.UUID, title, slug, description string) (*domain.Course, error) {
	if _, err := s.categories.FindByID(ctx, categoryID); err != nil {
		return nil, err
	}

	existing, err := s.courses.FindBySlug(ctx, slug)
	if err != nil && !errors.Is(err, domain.ErrCourseNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrCourseSlugAlreadyExists
	}

	course, err := domain.NewCourse(categoryID, title, slug, description)
	if err != nil {
		return nil, err
	}

	if err := s.courses.Create(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) Get(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	return s.courses.FindByID(ctx, id)
}

func (s *CourseService) List(ctx context.Context) ([]domain.Course, error) {
	return s.courses.List(ctx)
}

func (s *CourseService) Update(ctx context.Context, id uuid.UUID, title, description *string) (*domain.Course, error) {
	course, err := s.courses.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	newTitle := course.Title
	if title != nil {
		newTitle = *title
	}
	newDescription := course.Description
	if description != nil {
		newDescription = *description
	}

	if err := course.Update(newTitle, newDescription); err != nil {
		return nil, err
	}

	if err := s.courses.Update(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) ChangeThumbnail(ctx context.Context, id uuid.UUID, url string) (*domain.Course, error) {
	course, err := s.courses.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	course.ChangeThumbnail(url)
	if err := s.courses.Update(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.courses.Delete(ctx, id)
}

func (s *CourseService) Publish(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	course, err := s.courses.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := course.Publish(); err != nil {
		return nil, err
	}
	if err := s.courses.Update(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}

func (s *CourseService) Unpublish(ctx context.Context, id uuid.UUID) (*domain.Course, error) {
	course, err := s.courses.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := course.Unpublish(); err != nil {
		return nil, err
	}
	if err := s.courses.Update(ctx, course); err != nil {
		return nil, err
	}
	return course, nil
}
