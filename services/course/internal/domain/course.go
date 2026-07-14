package domain

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID          uuid.UUID
	CategoryID  uuid.UUID
	Title       string
	Slug        string
	Description string
	Thumbnail   string
	Published   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func NewCourse(categoryID uuid.UUID, title, slug, description string) (*Course, error) {
	if title == "" {
		return nil, ErrCourseTitleRequired
	}
	if slug == "" {
		return nil, ErrCourseSlugRequired
	}
	if categoryID == uuid.Nil {
		return nil, ErrCategoryNotFound
	}

	now := time.Now()
	return &Course{
		ID:          uuid.New(),
		CategoryID:  categoryID,
		Title:       title,
		Slug:        slug,
		Description: description,
		Published:   false,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (c *Course) Publish() error {
	if c.Published {
		return ErrCourseAlreadyPublished
	}
	c.Published = true
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Course) Unpublish() error {
	if !c.Published {
		return ErrCourseNotPublished
	}
	c.Published = false
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Course) Update(title, description string) error {
	if title == "" {
		return ErrCourseTitleRequired
	}
	c.Title = title
	c.Description = description
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Course) ChangeThumbnail(url string) {
	c.Thumbnail = url
	c.UpdatedAt = time.Now()
}

func (c *Course) IsDeleted() bool {
	return c.DeletedAt != nil
}

func (c *Course) Delete() {
	now := time.Now()
	c.DeletedAt = &now
}
