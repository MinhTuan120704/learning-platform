package domain

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID        uuid.UUID
	SectionID uuid.UUID
	Title     string
	Content   string
	VideoURL  string
	Duration  int // giây
	Position  int
	Published bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewLesson(sectionID uuid.UUID, title, content string, position int) (*Lesson, error) {
	if title == "" {
		return nil, ErrLessonTitleRequired
	}
	if sectionID == uuid.Nil {
		return nil, ErrSectionNotFound
	}

	now := time.Now()
	return &Lesson{
		ID:        uuid.New(),
		SectionID: sectionID,
		Title:     title,
		Content:   content,
		Position:  position,
		Published: false,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (l *Lesson) Publish() error {
	if l.Published {
		return ErrLessonAlreadyPublished
	}
	l.Published = true
	l.UpdatedAt = time.Now()
	return nil
}

func (l *Lesson) Unpublish() error {
	if !l.Published {
		return ErrLessonNotPublished
	}
	l.Published = false
	l.UpdatedAt = time.Now()
	return nil
}

func (l *Lesson) Move(position int) {
	l.Position = position
	l.UpdatedAt = time.Now()
}

func (l *Lesson) Update(title, content, videoURL string, duration int) error {
	if title == "" {
		return ErrLessonTitleRequired
	}
	l.Title = title
	l.Content = content
	l.VideoURL = videoURL
	l.Duration = duration
	l.UpdatedAt = time.Now()
	return nil
}
