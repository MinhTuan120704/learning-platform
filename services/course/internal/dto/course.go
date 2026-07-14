package dto

import "github.com/google/uuid"

type CreateCourseRequest struct {
	CategoryID  uuid.UUID `json:"category_id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
}

type UpdateCourseRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type CourseResponse struct {
	ID          uuid.UUID `json:"id"`
	CategoryID  uuid.UUID `json:"category_id"`
	Title       string    `json:"title"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"`
	Published   bool      `json:"published"`
}
