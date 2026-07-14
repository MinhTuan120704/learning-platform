package domain

import "errors"

var (
	ErrEnrollmentNotFound = errors.New("enrollment not found")
	ErrAlreadyEnrolled    = errors.New("already enrolled in this course")
	ErrUserIDRequired     = errors.New("user id is required")
	ErrCourseIDRequired   = errors.New("course id is required")
)
