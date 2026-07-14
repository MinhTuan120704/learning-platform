package validator

import (
	"errors"
	"regexp"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/dto"
)

var (
	ErrNameRequired     = errors.New("name is required")
	ErrEmailRequired    = errors.New("email is required")
	ErrEmailInvalid     = errors.New("email is invalid")
	ErrPasswordTooShort = errors.New("password must be at least 8 characters")
)

const minPasswordLen = 8

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func ValidateRegister(req dto.RegisterRequest) error {
	if req.Name == "" {
		return ErrNameRequired
	}

	if err := validateEmail(req.Email); err != nil {
		return err
	}

	if len(req.Password) < minPasswordLen {
		return ErrPasswordTooShort
	}

	return nil
}

func ValidateLogin(req dto.LoginRequest) error {
	if err := validateEmail(req.Email); err != nil {
		return err
	}

	if len(req.Password) < minPasswordLen {
		return ErrPasswordTooShort
	}

	return nil
}

func validateEmail(email string) error {
	if email == "" {
		return ErrEmailRequired
	}

	if !emailRegex.MatchString(email) {
		return ErrEmailInvalid
	}

	return nil
}

func ValidateUpdateUser(req dto.UpdateUserRequest) error {
	if req.Name != nil && *req.Name == "" {
		return ErrNameRequired
	}
	return nil
}
