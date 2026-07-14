package dto

import "github.com/google/uuid"

type CreateLessonRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content"`
	VideoURL string `json:"video_url"`
	Duration int    `json:"duration"`
}

type UpdateLessonRequest struct {
	Title    *string `json:"title"`
	Content  *string `json:"content"`
	VideoURL *string `json:"video_url"`
	Duration *int    `json:"duration"`
}

type LessonResponse struct {
	ID        uuid.UUID `json:"id"`
	SectionID uuid.UUID `json:"section_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	VideoURL  string    `json:"video_url"`
	Duration  int       `json:"duration"`
	Position  int       `json:"position"`
	Published bool      `json:"published"`
}
