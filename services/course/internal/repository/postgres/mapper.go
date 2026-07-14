package postgres

import (
	"time"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/domain"
	"github.com/google/uuid"
)

type categoryRow struct {
	ID          uuid.UUID
	Name        string
	Slug        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func mapCategory(row categoryRow) *domain.Category {
	return &domain.Category{
		ID:          row.ID,
		Name:        row.Name,
		Slug:        row.Slug,
		Description: row.Description,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

type courseRow struct {
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

func mapCourse(row courseRow) *domain.Course {
	return &domain.Course{
		ID:          row.ID,
		CategoryID:  row.CategoryID,
		Title:       row.Title,
		Slug:        row.Slug,
		Description: row.Description,
		Thumbnail:   row.Thumbnail,
		Published:   row.Published,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		DeletedAt:   row.DeletedAt,
	}
}

type sectionRow struct {
	ID          uuid.UUID
	CourseID    uuid.UUID
	Title       string
	Description string
	Position    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func mapSection(row sectionRow) *domain.Section {
	return &domain.Section{
		ID:          row.ID,
		CourseID:    row.CourseID,
		Title:       row.Title,
		Description: row.Description,
		Position:    row.Position,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

type lessonRow struct {
	ID        uuid.UUID
	SectionID uuid.UUID
	Title     string
	Content   string
	VideoURL  string
	Duration  int
	Position  int
	Published bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func mapLesson(row lessonRow) *domain.Lesson {
	return &domain.Lesson{
		ID:        row.ID,
		SectionID: row.SectionID,
		Title:     row.Title,
		Content:   row.Content,
		VideoURL:  row.VideoURL,
		Duration:  row.Duration,
		Position:  row.Position,
		Published: row.Published,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
