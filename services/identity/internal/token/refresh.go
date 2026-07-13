package token

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	goredis "github.com/redis/go-redis/v9"
)

var ErrRefreshTokenInvalid = errors.New("refresh token invalid or expired")

const refreshKeyPrefix = "refresh_token:"

type RefreshTokenService struct {
	rdb *goredis.Client
	ttl time.Duration
}

func NewRefreshTokenService(rdb *goredis.Client, ttl time.Duration) *RefreshTokenService {
	return &RefreshTokenService{rdb: rdb, ttl: ttl}
}

func (s *RefreshTokenService) Generate(ctx context.Context, userID uuid.UUID) (string, error) {
	raw := make([]byte, 32)

	if _, err := rand.Read(raw); err != nil {
		return "", err
	}
	plain := hex.EncodeToString(raw)

	key := refreshKeyPrefix + hashToken(plain)
	if err := s.rdb.Set(ctx, key, userID.String(), s.ttl).Err(); err != nil {
		return "", err
	}

	return plain, nil
}

func (s *RefreshTokenService) Validate(ctx context.Context, plain string) (uuid.UUID, error) {
	key := refreshKeyPrefix + hashToken(plain)

	val, err := s.rdb.Get(ctx, key).Result()
	if errors.Is(err, goredis.Nil) {
		return uuid.Nil, ErrRefreshTokenInvalid
	}
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := uuid.Parse(val)
	if err != nil {
		return uuid.Nil, ErrRefreshTokenInvalid
	}
	return userID, nil
}

func (s *RefreshTokenService) Rotate(ctx context.Context, oldPlain string) (string, uuid.UUID, error) {
	userID, err := s.Validate(ctx, oldPlain)
	if err != nil {
		return "", uuid.Nil, err
	}

	if err := s.Revoke(ctx, oldPlain); err != nil {
		return "", uuid.Nil, err
	}

	newPlain, err := s.Generate(ctx, userID)
	if err != nil {
		return "", uuid.Nil, err
	}

	return newPlain, userID, nil
}

func (s *RefreshTokenService) Revoke(ctx context.Context, plain string) error {
	key := refreshKeyPrefix + hashToken(plain)
	return s.rdb.Del(ctx, key).Err()
}

func hashToken(plain string) string {
	sum := sha256.Sum256([]byte(plain))
	return fmt.Sprintf("%x", sum)
}
