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

func NewUser(name string, email string, passwordHash string) (*User, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if passwordHash == "" {
		return nil, ErrInvalidPassword
	}

	now := time.Now()
	return &User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
