package service

import (
	"context"
	"errors"
	"time"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/domain"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/repository"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/security"
	"github.com/MinhTuan120704/learning-platform/services/identity/internal/token"
)

const defaultRoleName = "student"

type AuthService struct {
	users    repository.UserRepository
	roles    repository.RoleRepository
	password *security.PasswordService
	jwt      *token.JWTService
	refresh  *token.RefreshTokenService
}

func NewAuthService(users repository.UserRepository,
	roles repository.RoleRepository,
	password *security.PasswordService,
	jwt *token.JWTService,
	refresh *token.RefreshTokenService,
) *AuthService {
	return &AuthService{
		users:    users,
		roles:    roles,
		password: password,
		jwt:      jwt,
		refresh:  refresh,
	}
}

func (s *AuthService) Register(ctx context.Context, name, email, plainPassword string) (*domain.User, error) {
	existing, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		if !errors.Is(err, domain.ErrUserNotFound) {
			return nil, err
		}
	}
	if existing != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	hashPassword, err := s.password.Hash(plainPassword)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(name, email, hashPassword)
	if err != nil {
		return nil, err
	}

	if err := s.users.Create(ctx, user); err != nil {
		return nil, err
	}

	defaultRole, err := s.roles.FindByName(ctx, defaultRoleName)
	if err != nil {
		return nil, err
	}

	if err := s.users.AssignRole(ctx, user.ID, defaultRole.ID); err != nil {
		return nil, err
	}

	user.Roles = []domain.Role{*defaultRole}
	return user, nil
}

func (s *AuthService) Login(ctx context.Context, email, plainPassword string) (accessToken string, refreshToken string, expiresAt time.Time, err error) {
	user, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return "", "", time.Time{}, domain.ErrUserNotFound
		}
		return "", "", time.Time{}, err
	}

	if !s.password.Compare(user.PasswordHash, plainPassword) {
		return "", "", time.Time{}, domain.ErrInvalidCredentials
	}

	accessToken, expiresAt, err = s.jwt.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", time.Time{}, err
	}

	refreshToken, err = s.refresh.Generate(ctx, user.ID)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken, refreshToken, expiresAt, nil
}

func (s *AuthService) Refresh(ctx context.Context, oldRefreshToken string) (accessToken string, newRefreshToken string, expiresAt time.Time, err error) {
	newRefreshToken, userID, err := s.refresh.Rotate(ctx, oldRefreshToken)
	if err != nil {
		return "", "", time.Time{}, err
	}

	accessToken, expiresAt, err = s.jwt.GenerateAccessToken(userID)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken, newRefreshToken, expiresAt, nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {
	return s.refresh.Revoke(ctx, refreshToken)
}
