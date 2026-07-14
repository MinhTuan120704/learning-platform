package dto

import "github.com/google/uuid"

type CreateSectionRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateSectionRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type SectionResponse struct {
	ID          uuid.UUID `json:"id"`
	CourseID    uuid.UUID `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Position    int       `json:"position"`
}
