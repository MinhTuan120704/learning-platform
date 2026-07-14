package domain

import (
	"time"

	"github.com/google/uuid"
)

type Section struct {
	ID          uuid.UUID
	CourseID    uuid.UUID
	Title       string
	Description string
	Position    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewSection(courseID uuid.UUID, title, description string, position int) (*Section, error) {
	if title == "" {
		return nil, ErrSectionTitleRequired
	}
	if courseID == uuid.Nil {
		return nil, ErrCourseNotFound
	}

	now := time.Now()
	return &Section{
		ID:          uuid.New(),
		CourseID:    courseID,
		Title:       title,
		Description: description,
		Position:    position,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func (s *Section) Rename(title string) error {
	if title == "" {
		return ErrSectionTitleRequired
	}
	s.Title = title
	s.UpdatedAt = time.Now()
	return nil
}

func (s *Section) Move(position int) {
	s.Position = position
	s.UpdatedAt = time.Now()
}

func (s *Section) UpdateDescription(description string) {
	s.Description = description
	s.UpdatedAt = time.Now()
}
