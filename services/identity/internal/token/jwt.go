package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token expired")
)

type JWTService struct {
	secret    []byte
	accessTTL time.Duration
	issuer    string
}

func NewJWTService(secret string, accessTTL time.Duration, issuer string) *JWTService {
	return &JWTService{
		secret:    []byte(secret),
		accessTTL: accessTTL,
		issuer:    issuer,
	}
}

func (s *JWTService) GenerateAccessToken(userID uuid.UUID) (string, time.Time, error) {
	now := time.Now()
	expiresAt := now.Add(s.accessTTL)

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			Subject:   userID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := t.SignedString(s.secret)
	if err != nil {
		return "", time.Time{}, err
	}
	return signed, expiresAt, nil
}

func (s *JWTService) ParseAccessToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	parsed, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return s.secret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	if !parsed.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
