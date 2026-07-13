package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/MinhTuan120704/learning-platform/services/identity/internal/token"
)

const ContextUserIDKey = "userID"

func Authenticate(jwtSvc *token.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			return
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		claims, err := jwtSvc.ParseAccessToken(parts[1])
		if err != nil {
			status := http.StatusUnauthorized
			msg := "invalid token"
			if errors.Is(err, token.ErrExpiredToken) {
				msg = "token expired"
			}
			c.AbortWithStatusJSON(status, gin.H{"error": msg})
			return
		}

		c.Set(ContextUserIDKey, claims.UserID)
		c.Next()
	}
}

func UserIDFromContext(c *gin.Context) (uuid.UUID, bool) {
	val, exists := c.Get(ContextUserIDKey)
	if !exists {
		return uuid.Nil, false
	}
	id, ok := val.(uuid.UUID)
	return id, ok
}
