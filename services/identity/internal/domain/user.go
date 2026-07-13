package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID
	Name            string
	Email           string
	PasswordHash    string
	EmailVerifiedAt *time.Time
	Roles           []Role
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
