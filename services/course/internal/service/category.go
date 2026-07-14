package service

import (
	"context"
	"errors"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/repository"
	"github.com/google/uuid"
)

type CategoryService struct {
	categories repository.CategoryRepository
}

func NewCategoryService(categories repository.CategoryRepository) *CategoryService {
	return &CategoryService{categories: categories}
}

func (s *CategoryService) Create(ctx context.Context, name, slug, description string) (*domain.Category, error) {
	existing, err := s.categories.FindBySlug(ctx, slug)
	if err != nil && !errors.Is(err, domain.ErrCategoryNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrCategorySlugAlreadyExists
	}

	category, err := domain.NewCategory(name, slug, description)
	if err != nil {
		return nil, err
	}

	if err := s.categories.Create(ctx, category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) Get(ctx context.Context, id uuid.UUID) (*domain.Category, error) {
	return s.categories.FindByID(ctx, id)
}

func (s *CategoryService) List(ctx context.Context) ([]domain.Category, error) {
	return s.categories.List(ctx)
}

func (s *CategoryService) Update(ctx context.Context, id uuid.UUID, name, slug *string, description *string) (*domain.Category, error) {
	category, err := s.categories.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	newName := category.Name
	if name != nil {
		newName = *name
	}
	newSlug := category.Slug
	if slug != nil {
		newSlug = *slug
	}

	if newSlug != category.Slug {
		existing, err := s.categories.FindBySlug(ctx, newSlug)
		if err != nil && !errors.Is(err, domain.ErrCategoryNotFound) {
			return nil, err
		}
		if existing != nil {
			return nil, domain.ErrCategorySlugAlreadyExists
		}
	}

	if err := category.Rename(newName, newSlug); err != nil {
		return nil, err
	}
	if description != nil {
		category.UpdateDescription(*description)
	}

	if err := s.categories.Update(ctx, category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *CategoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.categories.Delete(ctx, id)
}
