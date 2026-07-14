package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type PermissionCache struct {
	client *redis.Client
	ttl    time.Duration
}

type CachedPermissions struct {
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

func NewPermissionCache(client *redis.Client, ttl time.Duration) *PermissionCache {
	return &PermissionCache{client: client, ttl: ttl}
}

func (c *PermissionCache) Get(ctx context.Context, userID string) (*CachedPermissions, error) {
	val, err := c.client.Get(ctx, "perms:"+userID).Result()
	if err == redis.Nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var perms CachedPermissions
	if err := json.Unmarshal([]byte(val), &perms); err != nil {
		return nil, err
	}
	return &perms, nil
}

func (c *PermissionCache) Set(ctx context.Context, userID string, perms CachedPermissions) error {
	data, err := json.Marshal(perms)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, "perms:"+userID, data, c.ttl).Err()
}
