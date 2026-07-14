package domain

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID          uuid.UUID
	Name        string
	Slug        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCategory(name, slug, description string) (*Category, error) {
	if name == "" {
		return nil, ErrCategoryNameRequired
	}
	if slug == "" {
		return nil, ErrCategorySlugRequired
	}

	now := time.Now()
	return &Category{
		ID:          uuid.New(),
		Name:        name,
		Slug:        slug,
		Description: description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (c *Category) Rename(name, slug string) error {
	if name == "" {
		return ErrCategoryNameRequired
	}
	if slug == "" {
		return ErrCategorySlugRequired
	}
	c.Name = name
	c.Slug = slug
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Category) UpdateDescription(description string) {
	c.Description = description
	c.UpdatedAt = time.Now()
}
