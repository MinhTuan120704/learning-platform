package domain

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID uuid.UUID

	Code string

	Description string

	CreatedAt time.Time

	UpdatedAt time.Time
}
