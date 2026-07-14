package domain

import "errors"

var (
	// Not found
	ErrCourseNotFound   = errors.New("course not found")
	ErrCategoryNotFound = errors.New("category not found")
	ErrSectionNotFound  = errors.New("section not found")
	ErrLessonNotFound   = errors.New("lesson not found")

	// Validation
	ErrCourseTitleRequired  = errors.New("course title is required")
	ErrCourseSlugRequired   = errors.New("course slug is required")
	ErrCategoryNameRequired = errors.New("category name is required")
	ErrCategorySlugRequired = errors.New("category slug is required")
	ErrSectionTitleRequired = errors.New("section title is required")
	ErrLessonTitleRequired  = errors.New("lesson title is required")

	// State
	ErrCourseAlreadyPublished = errors.New("course already published")
	ErrCourseNotPublished     = errors.New("course not published")
	ErrLessonAlreadyPublished = errors.New("lesson already published")
	ErrLessonNotPublished     = errors.New("lesson not published")

	// Conflict
	ErrCourseSlugAlreadyExists   = errors.New("course slug already exists")
	ErrCategorySlugAlreadyExists = errors.New("category slug already exists")

	// Invalid
	ErrInvalidPosition = errors.New("position must be >= 0")
	ErrInvalidDuration = errors.New("duration must be >= 0")
)
