package service

import (
	"context"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/repository"
	"github.com/google/uuid"
)

type LessonService struct {
	lessons  repository.LessonRepository
	sections repository.SectionRepository
}

func NewLessonService(lessons repository.LessonRepository, sections repository.SectionRepository) *LessonService {
	return &LessonService{lessons: lessons, sections: sections}
}

func (s *LessonService) Create(ctx context.Context, sectionID uuid.UUID, title, content, videoURL string, duration int) (*domain.Lesson, error) {
	if _, err := s.sections.FindByID(ctx, sectionID); err != nil {
		return nil, err
	}

	existing, err := s.lessons.ListBySectionID(ctx, sectionID)
	if err != nil {
		return nil, err
	}
	position := len(existing) + 1

	lesson, err := domain.NewLesson(sectionID, title, content, position)
	if err != nil {
		return nil, err
	}
	lesson.VideoURL = videoURL
	lesson.Duration = duration

	if err := s.lessons.Create(ctx, lesson); err != nil {
		return nil, err
	}
	return lesson, nil
}

func (s *LessonService) Get(ctx context.Context, id uuid.UUID) (*domain.Lesson, error) {
	return s.lessons.FindByID(ctx, id)
}

func (s *LessonService) ListBySection(ctx context.Context, sectionID uuid.UUID) ([]domain.Lesson, error) {
	return s.lessons.ListBySectionID(ctx, sectionID)
}

func (s *LessonService) Update(ctx context.Context, id uuid.UUID, title, content, videoURL *string, duration *int) (*domain.Lesson, error) {
	lesson, err := s.lessons.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	newTitle, newContent, newVideoURL, newDuration := lesson.Title, lesson.Content, lesson.VideoURL, lesson.Duration
	if title != nil {
		newTitle = *title
	}
	if content != nil {
		newContent = *content
	}
	if videoURL != nil {
		newVideoURL = *videoURL
	}
	if duration != nil {
		newDuration = *duration
	}

	if err := lesson.Update(newTitle, newContent, newVideoURL, newDuration); err != nil {
		return nil, err
	}
	if err := s.lessons.Update(ctx, lesson); err != nil {
		return nil, err
	}
	return lesson, nil
}

func (s *LessonService) Move(ctx context.Context, id uuid.UUID, position int) (*domain.Lesson, error) {
	lesson, err := s.lessons.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	lesson.Move(position)
	if err := s.lessons.Update(ctx, lesson); err != nil {
		return nil, err
	}
	return lesson, nil
}

func (s *LessonService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.lessons.Delete(ctx, id)
}

func (s *LessonService) Publish(ctx context.Context, id uuid.UUID) (*domain.Lesson, error) {
	lesson, err := s.lessons.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := lesson.Publish(); err != nil {
		return nil, err
	}
	if err := s.lessons.Update(ctx, lesson); err != nil {
		return nil, err
	}
	return lesson, nil
}

func (s *LessonService) Unpublish(ctx context.Context, id uuid.UUID) (*domain.Lesson, error) {
	lesson, err := s.lessons.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := lesson.Unpublish(); err != nil {
		return nil, err
	}
	if err := s.lessons.Update(ctx, lesson); err != nil {
		return nil, err
	}
	return lesson, nil
}
