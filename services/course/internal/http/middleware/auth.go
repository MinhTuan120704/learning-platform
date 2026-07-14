package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MinhTuan120704/learning-platform/services/course/internal/cache"
	"github.com/MinhTuan120704/learning-platform/services/course/internal/client"
)

func RequirePermission(permCache *cache.PermissionCache, identityClient *client.IdentityClient, permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing user identity"})
			return
		}

		ctx := c.Request.Context()

		cached, err := permCache.Get(ctx, userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "cache error"})
			return
		}

		var perms []string
		if cached != nil {
			perms = cached.Permissions
		} else {
			result, err := identityClient.GetUserPermissions(ctx, userID)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "identity service unavailable"})
				return
			}
			perms = result.Permissions
			_ = permCache.Set(ctx, userID, cache.CachedPermissions{Roles: result.Roles, Permissions: result.Permissions})
		}

		allowed := false
		for _, p := range perms {
			if p == permission {
				allowed = true
				break
			}
		}
		if !allowed {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
