package dto

import (
	"time"

	"github.com/google/uuid"
)

type GetUserResponse struct {
	ID              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at"`
}

type UpdateUserRequest struct {
	Name *string `json:"name"`
}
